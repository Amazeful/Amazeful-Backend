package server

import (
	"net/http"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

var (
	reqTimeout   = 2 * time.Minute
	requestLimit = 10
	limitTimeout = 10 * time.Second
)

type server struct {
	r *chi.Mux
}

func NewServer() *server {
	return &server{
		r: chi.NewRouter(),
	}
}

func (s *server) InitServer() error {
	var err error
	config := config.GetConfig()
	s.addDefaultMiddlewares()
	s.addRoutes()
	if config.TLS {
		err = http.ListenAndServeTLS(config.IpAddress+":"+config.Port, config.CertPath, config.KeyPath, s.r)
	} else {
		err = http.ListenAndServe(config.IpAddress+":"+config.Port, s.r)
	}
	if err != nil {
		return err
	}

	return nil
}

func (s *server) addDefaultMiddlewares() {
	s.r.Use(middleware.RequestID)
	s.r.Use(middleware.RealIP)
	s.r.Use(middleware.Logger)
	s.r.Use(middleware.Recoverer)
	s.r.Use(middleware.Timeout(reqTimeout))
	s.r.Use(httprate.Limit(requestLimit, limitTimeout, httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint)))
}
