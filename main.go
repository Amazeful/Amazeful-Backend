package main

import (
	"context"
	"log"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/server"
	"github.com/Amazeful/Amazeful-Backend/util"
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

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	//setup database
	zap.L().Info("starting database")
	err = util.InitDB(ctx)
	if err != nil {
		zap.L().Fatal("failed to init db", zap.Error(err))
	}
	client := util.GetMongoClient()
	defer client.Disconnect(ctx)

	//setup redis
	err = util.InitRedisClient(ctx)
	if err != nil {
		zap.L().Fatal("failed to init redis", zap.Error(err))
	}

	//setup server
	zap.L().Info("starting server")
	server := server.NewServer()
	err = server.InitServer()
	if err != nil {
		zap.L().Fatal("failed to init server", zap.Error(err))
	}
}
