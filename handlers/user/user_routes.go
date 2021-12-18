package user

import "github.com/go-chi/chi/v5"

func ProcessRoutes(r chi.Router) {
	r.Get("/", HandleGetUser)
	r.Post("/", HandleCreateUser)
}
