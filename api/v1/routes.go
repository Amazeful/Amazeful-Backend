package v1

import (
	"github.com/Amazeful/Amazeful-Backend/api/v1/channel"
	"github.com/Amazeful/Amazeful-Backend/util"

	"github.com/go-chi/chi/v5"
)

type V1 struct {
	*util.Resources
}

func NewV1(resources *util.Resources) *V1 {
	return &V1{
		Resources: resources,
	}
}
func (v *V1) ProcessRoutes(r chi.Router) {
	channelHandlers := channel.NewChannelHandler(v.Resources)

	r.Route("/channel", channelHandlers.ProcessRoutes)

}
