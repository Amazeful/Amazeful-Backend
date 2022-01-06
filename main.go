package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Amazeful/Amazeful-Backend/api/auth"
	v1 "github.com/Amazeful/Amazeful-Backend/api/v1"
	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/dataful"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"go.uber.org/zap"
)

var (
	reqTimeout   = 2 * time.Minute
	requestLimit = 10
	limitTimeout = 10 * time.Second
)

func main() {
	//setup the logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("Failed to init logger")
	}

	//load the config
	logger.Info("setting up config")
	config, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	//setup database
	logger.Info("starting database")
	mongo, err := dataful.NewMongoDB(ctx, config.ServerConfig.MongoURI)
	if err != nil {
		logger.Fatal("failed to init db", zap.Error(err))
	}
	defer mongo.Disconnect(ctx)

	redis, err := dataful.NewRedis(ctx, config.ServerConfig.RedisURI, config.ServerConfig.RedisPassword)
	if err != nil {
		logger.Fatal("failed to init cache", zap.Error(err))
	}

	twitchAPI := dataful.NewHelix(config.TwitchConfig.ClientID, config.TwitchConfig.ClientSecret)

	resources := &util.Resources{
		DB:        mongo,
		Cache:     redis,
		Logger:    logger,
		Config:    config,
		TwitchAPI: twitchAPI,
	}

	//setup server
	zap.L().Info("starting server")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(reqTimeout))
	r.Use(httprate.Limit(requestLimit, limitTimeout, httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint)))

	authHandler := auth.NewAuthHandler(resources)

	r.Route("/auth", authHandler.ProcessRoutes)
	r.Route("/v1", v1.SetupRoutes(resources))

	if config.ServerConfig.TLS {
		err = http.ListenAndServeTLS(config.ServerConfig.IpAddress+":"+config.ServerConfig.Port, config.ServerConfig.CertPath, config.ServerConfig.KeyPath, r)
	} else {
		err = http.ListenAndServe(config.ServerConfig.IpAddress+":"+config.ServerConfig.Port, r)
	}
	if err != nil {
		logger.Fatal("failed to init server", zap.Error(err))
	}

}
