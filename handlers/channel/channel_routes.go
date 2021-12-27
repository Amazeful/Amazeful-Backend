package channel

import (
	"github.com/Amazeful/Amazeful-Backend/middlewares"
	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {
	r.Use(middlewares.ChannelCtx)

	r.Get("/", HandleGetChannel)
}
