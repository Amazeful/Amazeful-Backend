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
	err := util.InitLogger()
	if err != nil {
		log.Fatal("Failed to init logger")
	}
	logger := util.GetLogger()
	//load the config
	logger.Info("setting up config")
	err = config.LoadConfig()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}
	cfg := config.GetConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()

	//setup database
	logger.Info("starting database")
	err = util.InitDB(ctx)
	if err != nil {
		logger.Fatal("failed to init db", zap.Error(err))
	}

	db := util.GetDB()
	defer db.Disconnect(context.Background())

	err = util.InitCache(ctx)
	if err != nil {
		logger.Fatal("failed to init cache", zap.Error(err))
	}

	//setup server
	logger.Info("starting server")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(reqTimeout))
	r.Use(httprate.Limit(requestLimit, limitTimeout, httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint)))

	r.Route("/auth", auth.ProcessRoutes)
	r.Route("/v1", v1.ProcessRoutes)

	if cfg.ServerConfig.TLS {
		err = http.ListenAndServeTLS(cfg.ServerConfig.IpAddress+":"+cfg.ServerConfig.Port, cfg.ServerConfig.CertPath, cfg.ServerConfig.KeyPath, r)
	} else {
		err = http.ListenAndServe(cfg.ServerConfig.IpAddress+":"+cfg.ServerConfig.Port, r)
	}
	if err != nil {
		logger.Fatal("failed to init server", zap.Error(err))
	}

}
