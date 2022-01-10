package command

import (
	"github.com/Amazeful/Amazeful-Backend/middlewares"
	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {
	r.Route("/{commandId}", func(r chi.Router) {
		r.Use(middlewares.CommandFromId)
		r.Get("/", HandleGetCommand)
		r.Patch("/", HandleUpdateCommand)
		r.Delete("/", HandleDeleteCommand)
	})
	r.Put("/", HandleCreateCommand)
}
