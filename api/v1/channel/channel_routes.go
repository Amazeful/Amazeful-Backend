package channel

import (
	"github.com/Amazeful/Amazeful-Backend/api/auth"
	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {

	r.Group(func(r chi.Router) {
		//Dashboard routes
		r.Use(auth.Authenticator)
		r.Use(ChannelFromSession)
		r.Get("/", HandleGetChannel)
	})

}
