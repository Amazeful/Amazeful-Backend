package user

import (
	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {
	r.Use(UserFromSession)
	r.Get("/", HandleGetUser)
}
