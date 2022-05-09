// go:build wireinject
//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/rest"
	"github.com/google/wire"
)

func initializeServer(ctx context.Context) *rest.Server {
	wire.Build(config.NewServerConfig)
	return NewServer()
}
