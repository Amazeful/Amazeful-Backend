package auth

import (
	"context"
	"net/http"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/consts"
	"github.com/Amazeful/Amazeful-Backend/models"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/lestrrat-go/jwx/jwa"
)

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
		jwt := models.NewJWT([]byte(config.GetConfig().JwtSignKey), jwa.HS256)
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
