package channel

import "github.com/go-chi/chi/v5"

func ChannelRoutes(r chi.Router) {
	r.Get("/", HandleGetChannel)
	r.Post("/", HandleCreateChannel)
}
