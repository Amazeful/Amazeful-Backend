package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/handlers/channel"
	"github.com/Amazeful/Amazeful-Backend/handlers/user"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func main() {

	//setup the logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to init logger")
	}
	defer logger.Sync()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	//load the config
	zap.L().Info("setting up config")
	err = config.LoadConfig()
	if err != nil {
		zap.L().Fatal("failed to load config", zap.Error(err))
	}

	//setup database
	zap.L().Info("starting database")
	err = util.InitDB()
	if err != nil {
		zap.L().Fatal("failed to init db", zap.Error(err))
	}
	client := util.GetDB()
	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(), nil)
	if err != nil {
		zap.L().Fatal("connection to db failed", zap.Error(err))
	}

	//setup server
	zap.L().Info("starting server")
	config := config.GetConfig()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	//setup routes
	r.Route("/channel", channel.ProcessRoutes)
	r.Route("/user", user.ProcessRoutes)

	if config.TLS {
		err = http.ListenAndServeTLS(config.IpAddress+":"+config.Port, config.CertPath, config.KeyPath, r)
	} else {
		err = http.ListenAndServe(config.IpAddress+":"+config.Port, r)
	}

	if err != nil {
		zap.L().Fatal("failed to start server", zap.Error(err))
	}

}
