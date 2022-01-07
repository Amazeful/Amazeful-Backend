package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/util"

	"github.com/Amazeful/dataful"
	"github.com/Amazeful/dataful/models"
	"github.com/Amazeful/helix"
	"github.com/lestrrat-go/jwx/jwa"
)

const JWTCookieName = "amazing_token"

func HandleTwitchLogin(rw http.ResponseWriter, req *http.Request) {
	cfg := config.GetConfig()
	http.Redirect(rw, req, cfg.GetTwitchOauthConfig().AuthCodeURL(cfg.TwitchConfig.State), http.StatusTemporaryRedirect)
}

func HandleTwitchCallback(rw http.ResponseWriter, req *http.Request) {
	cfg := config.GetConfig()

	state := req.URL.Query().Get("state")
	code := req.URL.Query().Get("code")

	//Check state value
	if cfg.TwitchConfig.State != state {
		err := fmt.Errorf("invalid state value received -- %s", state)
		util.WriteError(rw, err, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	//Use the code to get tokens from twitch
	token, err := cfg.GetTwitchOauthConfig().Exchange(req.Context(), code)
	if err != nil {
		util.WriteError(rw, err, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	//Create api client
	client, err := helix.NewClient(&helix.Options{ClientID: cfg.TwitchConfig.ClientID, UserAccessToken: token.AccessToken})
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	//get user from twitch
	twitchUser, err := client.GetMe()
	if err != nil || twitchUser.ResponseCommon.StatusCode != http.StatusOK {
		util.WriteError(rw, err, http.StatusInternalServerError, "Failed to get user data from Twitch.")
		return
	}

	//get channel from twitch
	twitchChannel, err := client.GetChannelInformationById(twitchUser.Data.ID)
	if err != nil || twitchChannel.ResponseCommon.StatusCode != http.StatusOK {
		util.WriteError(rw, err, http.StatusInternalServerError, "Failed to get channel data from Twitch.")
		return
	}

	ru := util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionUser)
	rc := util.GetDB().Repository(dataful.DBAmazeful, dataful.CollectionChannel)

	channel := models.NewChannel(rc)
	err = channel.FindByChannelId(req.Context(), twitchChannel.Data.BroadcasterID)
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	channel.ChannelId = twitchChannel.Data.BroadcasterID
	channel.BroadcasterName = twitchChannel.Data.BroadcasterName
	channel.Language = twitchChannel.Data.BroadcasterLanguage
	channel.GameId = twitchChannel.Data.GameID
	channel.GameName = twitchChannel.Data.GameName
	channel.Title = twitchChannel.Data.Title

	if channel.Loaded() {
		err = channel.Update(req.Context())
	} else {
		err = channel.Create(req.Context())
	}
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	//make a new user using tokens
	user := models.NewUser(ru)
	err = user.FindByUserId(req.Context(), twitchUser.Data.ID)
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	user.UserID = twitchUser.Data.ID
	user.Login = twitchUser.Data.Login
	user.DisplayName = twitchUser.Data.DisplayName
	user.Type = twitchUser.Data.Type
	user.BroadcasterType = twitchUser.Data.BroadcasterType
	user.Description = twitchUser.Data.Description
	user.ProfileImageURL = twitchUser.Data.ProfileImageURL
	user.OfflineImageURL = twitchUser.Data.OfflineImageURL
	user.ViewCount = twitchUser.Data.ViewCount
	user.AccessToken = token.AccessToken
	user.RefreshToken = token.RefreshToken
	user.Channel = channel.ID

	if user.Loaded() {
		err = user.Update(req.Context())
	} else {
		err = user.Create(req.Context())
	}
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if user.Suspended {
		err = fmt.Errorf("user is suspended id: %s, name: %s", user.UserID, user.Login)
		util.WriteError(rw, err, http.StatusUnauthorized, "Your account has been suspended.")
		return
	}

	//session and token expires in a day
	expiry := time.Now().Add(time.Hour * 24)

	//Make a new session for user
	session := models.NewSession(util.GetCache())
	session.GenerateSessionId()
	session.User = user.ID
	session.SelectedChannel = channel.ID

	//Make a new jwt for token
	jwt := models.NewJWT([]byte(cfg.ServerConfig.JwtSignKey), jwa.HS256)
	tokenString, err := jwt.Encode(session.SessionId, expiry)
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	//set the new session
	err = session.SetSession(req.Context(), time.Until(expiry))
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	//Add in the token to cookie
	http.SetCookie(rw, &http.Cookie{
		Name:     JWTCookieName,
		Value:    tokenString,
		Expires:  expiry,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	})
}
