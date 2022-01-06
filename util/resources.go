package util

import (
	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/dataful"
	"go.uber.org/zap"
)

type Resources struct {
	DB        dataful.Database
	Cache     dataful.Cache
	Logger    *zap.Logger
	Config    *config.Config
	TwitchAPI dataful.TwitchAPI
}
