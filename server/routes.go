package server

import (
	"github.com/Amazeful/Amazeful-Backend/handlers/auth"
	"github.com/Amazeful/Amazeful-Backend/handlers/channel"
	"github.com/Amazeful/Amazeful-Backend/handlers/user"
)

func (s *server) addRoutes() {
	s.r.Route("/auth", auth.ProcessRoutes)
	s.r.Route("/channel", channel.ProcessRoutes)
	s.r.Route("/user", user.ProcessRoutes)
}
