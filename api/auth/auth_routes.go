package auth

import (
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/go-chi/chi/v5"
)

type AuthHandler struct {
	*util.Resources
}

func NewAuthHandler(resources *util.Resources) *AuthHandler {
	return &AuthHandler{
		Resources: resources,
	}
}

func (ah *AuthHandler) ProcessRoutes(r chi.Router) {
	r.Route("/twitch", func(r chi.Router) {
		r.Get("/login", ah.HandleTwitchLogin)
		r.Get("/callback", ah.HandleTwitchCallback)
	})
}
