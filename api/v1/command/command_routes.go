package command

import "github.com/go-chi/chi/v5"

func ProcessRoutes(r chi.Router) {

	r.Route("/{commandId}", func(r chi.Router) {
		r.Get("/", HandleGetCommand)
		r.Post("/", HandleCreateCommand)
		r.Patch("/", HandleUpdateCommand)
		r.Delete("/", HandleDeleteCommand)
	})

}
