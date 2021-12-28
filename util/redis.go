package util

import (
	"context"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/go-redis/redis/v8"
)

type IRedis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type Redis struct {
	client *redis.Client
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}

func (r *Redis) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

var rc IRedis

//InitRedis initializes Redis client
func InitRedisClient(ctx context.Context) error {
	redis := redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().RedisURI,
		Password: config.GetConfig().RedisPassword,
		DB:       0,
	})
	rx := redis.Ping(ctx)
	if rx.Err() != nil {
		return rx.Err()
	}

	rc = &Redis{
		client: redis,
	}
	return nil
}

//GetRedis returns redis client
func GetRedis() IRedis {
	return rc
}

func SetRedis(redis IRedis) {
	rc = redis
}
