package channel

import (
	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {
	r.Route("/{channelId}", func(r chi.Router) {
		r.Get("/", HandleGetChannel)
		r.Patch("/", HandleUpdateChannel)
	})
}
