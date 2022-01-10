package v1

import (
	"github.com/Amazeful/Amazeful-Backend/api/v1/channel"
	"github.com/Amazeful/Amazeful-Backend/api/v1/command"

	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {
	r.Route("/channel", channel.ProcessRoutes)
	r.Route("/command", command.ProcessRoutes)
}
