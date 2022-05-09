package rest

import (
	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/dataful/cache"
	"github.com/Amazeful/dataful/db"
	"github.com/go-chi/chi/v5"
)

type RouterCommon struct {
	DB             db.DB
	Cache          cache.Cache
	ResponseWriter ResponseWriter
	ServerConfig   *config.ServerConfig
}

type Router interface {
	Process(r chi.Router)
}
