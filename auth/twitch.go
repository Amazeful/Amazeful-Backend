package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/rest"
	"github.com/Amazeful/dataful/common"
	"github.com/Amazeful/dataful/models"
	"github.com/Amazeful/dataful/store"
	"github.com/Amazeful/helix"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

type TwitchAuth struct {
	twitchConfig   *config.TwitchConfig
	responseWriter rest.ResponseWriter
	channelStore   store.ChannelStore
	userStore      store.UserStore
	authenticator  Authenticator
	oauth2         *oauth2.Config
}

func NewTwitchAuth(
	twitchConfig *config.TwitchConfig,
	serverConfig *config.ServerConfig,
	responseWriter rest.ResponseWriter,
	channelStore store.ChannelStore,
	userStore store.UserStore,
	authenticator Authenticator,
) *TwitchAuth {
	return &TwitchAuth{
		twitchConfig:   twitchConfig,
		responseWriter: responseWriter,
		channelStore:   channelStore,
		userStore:      userStore,
		authenticator:  authenticator,
		oauth2: &oauth2.Config{
			ClientID:     twitchConfig.ClientID,
			ClientSecret: twitchConfig.ClientSecret,
			Endpoint:     twitch.Endpoint,
			RedirectURL:  serverConfig.ServerURL + "/auth/twitch/callback",
			Scopes:       []string{"user:read:email"},
		},
	}
}

//HandleTwitchLogin redirects user to Twitch for open auth.
func (ta *TwitchAuth) HandleTwitchLogin(rw http.ResponseWriter, req *http.Request) {

	http.Redirect(rw, req, ta.oauth2.AuthCodeURL(ta.twitchConfig.State), http.StatusTemporaryRedirect)
}

//HandleTwitchCallback handles callback request from Twitch oauth.
func (ta *TwitchAuth) HandleTwitchCallback(rw http.ResponseWriter, req *http.Request) {
	state := req.URL.Query().Get("state")
	code := req.URL.Query().Get("code")

	//Check state value
	if !ta.isStateValid(state) {
		err := fmt.Errorf("invalid state value received -- %s", state)
		ta.responseWriter.WriteError(rw, http.StatusInternalServerError, err, rest.MessageInternalServerError)
		return
	}

	//Use the code to get tokens from twitch
	token, err := ta.oauth2.Exchange(req.Context(), code)
	if err != nil {
		ta.responseWriter.WriteError(rw, http.StatusUnauthorized, err, rest.MessageUnauthorized)
		return
	}

	//Create api client
	client, err := helix.NewClient(&helix.Options{ClientID: ta.twitchConfig.ClientID, UserAccessToken: token.AccessToken})
	if err != nil {
		ta.responseWriter.WriteError(rw, http.StatusInternalServerError, err, rest.MessageInternalServerError)
		return
	}

	user, err := ta.getUser(req.Context(), client)
	if err != nil {
		ta.responseWriter.WriteError(rw, http.StatusInternalServerError, err, rest.MessageInternalServerError)
		return
	}

	channel, err := ta.getChannel(req.Context(), client, user.UserID)
	if err != nil {
		ta.responseWriter.WriteError(rw, http.StatusInternalServerError, err, rest.MessageInternalServerError)
		return
	}

	err = ta.authenticator.CreateToken(req.Context(), rw, user.ID, channel.ID)
	if err != nil {
		ta.responseWriter.WriteError(rw, http.StatusInternalServerError, err, rest.MessageInternalServerError)
		return
	}

}

//isStateValid validates the state returned from twitch callback.
func (ta *TwitchAuth) isStateValid(state string) bool {
	return ta.twitchConfig.State == state
}

//getChannel gets channel from Twitch API and returns a channel object.
func (ta *TwitchAuth) getChannel(ctx context.Context, client *helix.Client, channelId string) (*models.Channel, error) {
	//get channel from twitch
	twitchChannel, err := client.GetChannelInformationById(channelId)
	if err != nil {
		return nil, err
	}
	if twitchChannel.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("twitch returned status %d", twitchChannel.StatusCode)
	}

	channel, err := ta.channelStore.GetByChanneId(ctx, twitchChannel.Data.BroadcasterID, common.TwitchPlatform)
	if err != nil {
		return nil, err
	}

	channel.ChannelId = channelId
	channel.BroadcasterName = twitchChannel.Data.BroadcasterName
	channel.Language = twitchChannel.Data.BroadcasterLanguage
	channel.GameName = twitchChannel.Data.GameName
	channel.Title = twitchChannel.Data.Title
	channel.Platform = common.TwitchPlatform

	if channel.IsLoaded() {
		err = ta.channelStore.Update(ctx, channel)
	} else {
		err = ta.channelStore.Create(ctx, channel)
	}

	if err != nil {
		return nil, err
	}

	return channel, nil
}

//getUser gets user from Twitch API and returns a user object.
func (ta *TwitchAuth) getUser(ctx context.Context, client *helix.Client) (*models.User, error) {
	//get channel from twitch
	twitchUser, err := client.GetMe()
	if err != nil {
		return nil, err
	}
	if twitchUser.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("twitch returned status %d", twitchUser.StatusCode)
	}

	user, err := ta.userStore.GetByUserId(ctx, twitchUser.Data.ID)
	if err != nil {
		return nil, err
	}

	user.UserID = twitchUser.Data.ID
	user.Login = twitchUser.Data.Login
	user.DisplayName = twitchUser.Data.DisplayName
	user.ProfileImageURL = twitchUser.Data.ProfileImageURL
	user.ViewCount = twitchUser.Data.ViewCount

	if user.IsLoaded() {
		err = ta.userStore.Update(ctx, user)
	} else {
		err = ta.userStore.Create(ctx, user)
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}
