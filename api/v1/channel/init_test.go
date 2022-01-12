package channel

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	godotenv.Load("../../../.env")
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config -- %v", err)
	}

	cfg := config.GetConfig()
	err = util.InitLogger()
	if err != nil {
		log.Fatalf("failed to init logger -- %v", err)
	}

	err = util.InitDB(context.Background(), cfg.ServerConfig.MongoURI)
	if err != nil {
		log.Fatalf("failed to init db -- %v", err)
	}

	err = util.InitCache(context.Background(), cfg.ServerConfig.RedisURI, cfg.ServerConfig.RedisPassword)
	if err != nil {
		log.Fatalf("failed to init cache -- %v", err)
	}

	os.Exit(m.Run())
}
