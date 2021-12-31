package server

import (
	"github.com/Amazeful/Amazeful-Backend/api/auth"
	v1 "github.com/Amazeful/Amazeful-Backend/api/v1"
)

func (s *server) addRoutes() {
	s.r.Route("/auth", auth.ProcessRoutes)
	s.r.Route("/v1", v1.ProcessRoutes)
}
