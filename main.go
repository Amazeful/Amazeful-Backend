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
	"github.com/joho/godotenv"
)

var (
	reqTimeout   = 2 * time.Minute
	requestLimit = 10
	limitTimeout = 10 * time.Second
	ctx, cancel  = context.WithTimeout(context.Background(), 5*time.Minute)
)

func main() {
	initConfig()
	initServices()
	defer cleanupServices()
	initServer()
}

func initConfig() {
	godotenv.Load()
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to init config -- %v", err)
	}
}

func initServices() {
	cfg := config.GetConfig()
	err := util.InitLogger()
	if err != nil {
		log.Fatalf("failed to init logger -- %v", err)
	}

	err = util.InitDB(ctx, cfg.ServerConfig.MongoURI)
	if err != nil {
		log.Fatalf("failed to init db -- %v", err)
	}

	err = util.InitCache(ctx, cfg.ServerConfig.RedisURI, cfg.ServerConfig.RedisPassword)
	if err != nil {
		log.Fatalf("failed to init cache -- %v", err)
	}
}

func initServer() {
	cfg := config.GetConfig()
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(reqTimeout))
	r.Use(httprate.Limit(requestLimit, limitTimeout, httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint)))

	//routers
	r.Route("/auth", auth.ProcessRoutes)
	r.Route("/v1", v1.ProcessRoutes)

	var err error
	if cfg.ServerConfig.TLS {
		err = http.ListenAndServeTLS(cfg.ServerConfig.IpAddress+":"+cfg.ServerConfig.Port, cfg.ServerConfig.CertPath, cfg.ServerConfig.KeyPath, r)
	} else {
		err = http.ListenAndServe(cfg.ServerConfig.IpAddress+":"+cfg.ServerConfig.Port, r)
	}
	if err != nil {
		log.Fatalf("failed to init server -- %v", err)
	}
}

func cleanupServices() {
	cancel()
	util.DB().Disconnect(context.Background())
}
