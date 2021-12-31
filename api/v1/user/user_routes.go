package user

import (
	"github.com/Amazeful/Amazeful-Backend/middlewares"
	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {
	r.Use(middlewares.UserFromSession)
	r.Get("/", HandleGetUser)
}
