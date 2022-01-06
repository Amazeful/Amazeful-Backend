package v1

import (
	"github.com/Amazeful/Amazeful-Backend/api/v1/channel"
	"github.com/Amazeful/Amazeful-Backend/util"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(resources *util.Resources) func(r chi.Router) {
	return func(r chi.Router) {
		channelHandlers := channel.NewChannelHandler(resources)
		r.Route("/channel", channelHandlers.ProcessRoutes)
	}
}
