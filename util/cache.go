package util

import (
	"context"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/dataful"
)

var cache dataful.Cache

//InitCache initializes the cache
func InitCache(ctx context.Context) error {
	config := config.GetConfig()
	client, err := dataful.NewRedis(ctx, config.ServerConfig.RedisURI, config.ServerConfig.RedisPassword)
	if err != nil {
		return err
	}

	cache = client
	return nil
}

//GetCache returns global cache
func GetCache() dataful.Cache {
	return cache
}
