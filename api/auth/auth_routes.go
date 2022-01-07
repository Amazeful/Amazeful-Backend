package auth

import (
	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {
	r.Route("/twitch", func(r chi.Router) {
		r.Get("/login", HandleTwitchLogin)
		r.Get("/callback", HandleTwitchCallback)
	})
}
