package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/util"
)

func HandleTwitchLogin(rw http.ResponseWriter, req *http.Request) {
	twitchConfig := config.GetTwitchConfig()
	http.Redirect(rw, req, twitchConfig.OauthConfig.AuthCodeURL(twitchConfig.State), http.StatusTemporaryRedirect)
}

func HandleTwitchCallback(rw http.ResponseWriter, req *http.Request) {
	twitchConfig := config.GetTwitchConfig()

	receivedState := req.URL.Query().Get("state")
	code := req.URL.Query().Get("code")

	if twitchConfig.State != receivedState {
		util.WriteError(rw, errors.New("invalid state value received"), http.StatusUnauthorized, consts.ErrStrUnauthorized)
		return
	}

	token, err := twitchConfig.OauthConfig.Exchange(req.Context(), code)

	if err != nil {
		util.WriteError(rw, err, http.StatusUnauthorized, consts.ErrStrUnauthorized)
		return
	}

	// jwt, err := config.EncodeJWT(token)

	log.Print(token)
	return
}
