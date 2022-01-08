package channel

import (
	"github.com/Amazeful/Amazeful-Backend/middlewares"
	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {
	r.Route("/{channelId}", func(r chi.Router) {
		r.Use(middlewares.ChannelFromId)
		r.Get("/", HandleGetChannel)
		r.Patch("/", HandleUpdateChannel)
		r.Get("/commands", HandleGetChannelCommands)
		r.Get("/filters", HandleGetChannelFilters)
		// r.Put("/new-command", HandleCreateCommand)
	})
}
