package channel

import (
	"github.com/Amazeful/Amazeful-Backend/middlewares"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/go-chi/chi/v5"
)

type ChannelApi struct {
	R util.Repository
}

func ProcessRoutes(r chi.Router) {

	r.Route("/{channelId}", func(r chi.Router) {
		r.Use(middlewares.ChannelFromId)
		r.Get("/", HandleGetChannel)
		r.Patch("/", HandleUpdateChannel)
	})

}
