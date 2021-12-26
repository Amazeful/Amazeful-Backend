package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util"
)

const JWTCookieName = "amazing_token"

func HandleTwitchLogin(rw http.ResponseWriter, req *http.Request) {
	twitchOauthConfig := config.GetTwitchOauthConfig()
	http.Redirect(rw, req, twitchOauthConfig.AuthCodeURL(config.GetTwitchConfig().State), http.StatusTemporaryRedirect)
}

func HandleTwitchCallback(rw http.ResponseWriter, req *http.Request) {
	twitchConfig := config.GetTwitchConfig()
	twitchOauthConfig := config.GetTwitchOauthConfig()

	//Get state and token value
	receivedState := req.URL.Query().Get("state")
	code := req.URL.Query().Get("code")

	//Check state value
	if twitchConfig.State != receivedState {
		util.WriteError(rw, errors.New("invalid state value received"), http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	//Use the code to get tokens from twitch
	token, err := twitchOauthConfig.Exchange(req.Context(), code)
	if err != nil {
		util.WriteError(rw, err, http.StatusUnauthorized, "Failed to get tokens from Twitch.")
		return
	}

	collection := util.GetCollection(consts.CollectionUser)

	//make a new user using tokens
	user := models.NewUser(collection)
	user.AccessToken = token.AccessToken
	user.RefreshToken = token.RefreshToken

	//get user from Twitch api, make or update user in db
	err = user.GetUserFromTwitch(req.Context())
	if err != nil {
		util.WriteError(rw, err, http.StatusInternalServerError, "Failed to get user data from Twitch.")
		return
	}

	//session and token expires in a day
	expiry := time.Now().Add(time.Hour * 24)

	//Make a new session for user
	session := models.NewSession(util.GetRedis())
	session.GenerateSessionId()
	session.User = user.ID.String()

	//Make a new jwt for token
	jwt := models.NewJWT()
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

//Authenticator middleware authenticates requests
func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//Get the token from cookie
		cookie, err := req.Cookie(JWTCookieName)
		if err != nil {
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		//Parse and validate the token
		jwt := models.NewJWT()
		t, err := jwt.Decode(cookie.Value)
		if err != nil {
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		//Get session id from token claims
		sid, ok := t.Get(models.SessionIdKey)
		if !ok {
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		session := models.NewSession(util.GetRedis())
		session.SessionId = sid.(string)
		err = session.GetSession(req.Context())
		if err != nil {
			http.Error(rw, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), consts.CtxSession, session)

		next.ServeHTTP(rw, req.WithContext(ctx))
	})
}
