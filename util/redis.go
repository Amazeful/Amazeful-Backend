package util

import (
	"context"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/go-redis/redis/v8"
)

var rc *redis.Client

//InitRedis initializes Redis client
func InitRedis(ctx context.Context) error {
	rc = redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().RedisURI,
		Password: config.GetConfig().RedisPassword,
		DB:       0,
	})
	rx := rc.Ping(ctx)
	return rx.Err()
}

//GetRedis returns redis client
func GetRedis() *redis.Client {
	return rc
}
