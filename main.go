package main

import (
	"context"
	"time"

	"github.com/Amazeful/dataful/db"
	mongostore "github.com/Amazeful/dataful/store/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var (
	reqTimeout   = 2 * time.Minute
	requestLimit = 10
	limitTimeout = 10 * time.Second
	ctx, cancel  = context.WithTimeout(context.Background(), 5*time.Minute)
)

func main() {

	logger := zap.NewProductionConfig()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	db := db.NewMongoDB()

	channelStore := mongostore.NewMongoChannelStore()
}

// func initConfig() {
// 	godotenv.Load()
// 	err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatalf("failed to init config -- %v", err)
// 	}
// }

// func initServices() {
// 	cfg := config.GetConfig()
// 	err := util.InitAllServices(ctx, cfg)
// 	if err != nil {
// 		log.Fatalf("failed to services -- %v", err)
// 	}
// }

// func initServer() {
// 	cfg := config.GetConfig()
// 	r := chi.NewRouter()
// 	r.Use(chimid.RequestID)
// 	r.Use(chimid.RealIP)
// 	r.Use(chimid.Logger)
// 	r.Use(chimid.Recoverer)
// 	r.Use(chimid.Timeout(reqTimeout))
// 	r.Use(httprate.Limit(requestLimit, limitTimeout, httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint)))

// 	//routers
// 	r.Route("/auth", auth.ProcessRoutes)
// 	r.Route("/v1", v1.ProcessRoutes)

// 	var err error
// 	if cfg.ServerConfig.TLS {
// 		err = http.ListenAndServeTLS(cfg.ServerConfig.IpAddress+":"+cfg.ServerConfig.Port, cfg.ServerConfig.CertPath, cfg.ServerConfig.KeyPath, r)
// 	} else {
// 		err = http.ListenAndServe(cfg.ServerConfig.IpAddress+":"+cfg.ServerConfig.Port, r)
// 	}
// 	if err != nil {
// 		log.Fatalf("failed to init server -- %v", err)
// 	}
// }

// func cleanupServices() {
// 	cancel()
// 	util.DB().Disconnect(context.Background())
// }
