package util

import (
	"context"

	"github.com/Amazeful/dataful"
)

var cache dataful.Cache

//InitCache initializes the cache
func InitCache(ctx context.Context, uri, password string) error {
	var err error
	cache, err = dataful.NewRedis(ctx, uri, password)
	return err
}

//Cache returns global cache
func Cache() dataful.Cache {
	return cache
}
