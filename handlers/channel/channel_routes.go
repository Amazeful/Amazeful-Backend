package channel

import (
	"github.com/Amazeful/Amazeful-Backend/handlers/auth"
	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {

	r.Group(func(r chi.Router) {
		//Dashboard routes
		r.Use(auth.Authenticator)
		r.Use(ChannelFromSession)
		r.Get("/selected", HandleGetChannel)
	})

	r.Route("/{channelName}", func(r chi.Router) {
		r.Use(ChannelFromParam)
	})

}
