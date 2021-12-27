package main

import (
	"github.com/Amazeful/Amazeful-Backend/handlers/channel"
	"github.com/Amazeful/Amazeful-Backend/handlers/user"
	"github.com/Amazeful/Amazeful-Backend/middlewares"
	"github.com/go-chi/chi/v5"
)

func ProcessRoutes(r chi.Router) {
	r.Use(middlewares.Authenticator)
	r.Route("/channel", channel.ProcessRoutes)
	r.Route("/user", user.ProcessRoutes)
}
