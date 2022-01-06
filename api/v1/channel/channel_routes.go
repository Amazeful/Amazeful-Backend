package channel

import (
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/go-chi/chi/v5"
)

type ChannelHandler struct {
	*util.Resources
}

func NewChannelHandler(resources *util.Resources) *ChannelHandler {
	return &ChannelHandler{
		Resources: resources,
	}
}

func (ch *ChannelHandler) ProcessRoutes(r chi.Router) {
	r.Route("/{channelId}", func(r chi.Router) {
		r.Use(ch.ChannelFromId)
		r.Get("/", ch.HandleGetChannel)
		r.Patch("/", ch.HandleUpdateChannel)
	})
}
