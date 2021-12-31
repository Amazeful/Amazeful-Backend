package util

import (
	"context"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/go-redis/redis/v8"
)

var rc Redis

type Redis interface {
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
}

type RedisClient struct {
	client *redis.Client
}

func (r *RedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}

func (r *RedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

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

	rc = &RedisClient{
		client: redis,
	}
	return nil
}

//GetRedis returns redis client
func GetRedis() Redis {
	return rc
}

func SetRedis(redis Redis) {
	rc = redis
}
