package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/util"
)

func HandleTwitchLogin(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, config.GetConfig().Twitch.OauthConfig.AuthCodeURL(config.GetConfig().Twitch.State), http.StatusTemporaryRedirect)
}

func HandleTwitchCallback(rw http.ResponseWriter, req *http.Request) {
	actualState := config.GetConfig().Twitch.State
	receivedState := req.URL.Query().Get("state")

	if actualState != receivedState {
		util.WriteError(rw, errors.New("invalid state value received"), http.StatusUnauthorized, "invalid state value received")
		return
	}

	code := req.URL.Query().Get("code")

	token, err := config.GetConfig().Twitch.OauthConfig.Exchange(req.Context(), code)
	if err != nil {
		util.WriteError(rw, err, http.StatusUnauthorized, "failed to get a token from twitch")
		return
	}

	log.Print(token)
	return
}
