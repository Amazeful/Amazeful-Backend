package rest

import (
	"net/http"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	chi "github.com/go-chi/chi/v5"
	chimid "github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	config               *config.ServerConfig
	protectedControllers []Controller
	reqTimeout           time.Duration
}

func NewServer(config *config.ServerConfig, protectedControllers []Controller) *Server {
	return &Server{
		protectedControllers: protectedControllers,
		reqTimeout:           3 * time.Minute,
	}
}

func (s *Server) Start() error {
	h := s.createHandler()

	var err error
	if s.config.TLS {
		err = http.ListenAndServeTLS(s.config.IpAddress+":"+s.config.Port, s.config.CertPath, s.config.KeyPath, h)
	} else {
		err = http.ListenAndServe(s.config.IpAddress+":"+s.config.Port, h)
	}

	return err
}

func (s *Server) createHandler() http.Handler {
	router := chi.NewRouter()
	router.Use(chimid.RequestID)
	router.Use(chimid.RealIP)
	router.Use(chimid.Logger)
	router.Use(chimid.Recoverer)
	router.Use(chimid.Timeout(s.reqTimeout))

	//add routes
	for _, c := range s.protectedControllers {
		router.Route(c.BasePath(), c.Routes)
	}

	return router
}
