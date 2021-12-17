package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to init logger")
	}
	defer logger.Sync()

	undo := zap.ReplaceGlobals(logger)
	defer undo()

	zap.L().Info("starting server")

	config := config.GetConfig()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	if config.CertPath != "" && config.KeyPath != "" {
		http.ListenAndServeTLS(config.IpAddress+":"+config.Port, config.CertPath, config.KeyPath, r)
	} else {
		http.ListenAndServe(config.IpAddress+":"+config.Port, r)
	}
}
