package authmodel

import (
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/go-chi/jwtauth/v5"
)

type JWT struct {
	jwtAuth *jwtauth.JWTAuth
}

func NewJWT() *JWT {
	return &JWT{
		jwtAuth: jwtauth.New("HS256", config.GetConfig().TokenSecret, nil),
	}
}

func (jwt *JWT) Encode(issuer string, userId int) (string, error) {
	claims := map[string]interface{}{}
	jwtauth.SetExpiryIn(claims, time.Hour*24)
	jwtauth.SetIssuedNow(claims)
	claims["issuer"] = issuer
	claims["user_id"] = userId

	_, jwtString, err := jwt.jwtAuth.Encode(claims)
	if err != nil {
		return "", err
	}

	return jwtString, nil
}
