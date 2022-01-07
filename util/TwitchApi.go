package util

import (
	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/dataful"
)

var twitchAPI dataful.TwitchAPI

func InitTwitchAPI() {
	cfg := config.GetConfig()
	twitchAPI = dataful.NewHelix(cfg.TwitchConfig.ClientID, cfg.TwitchConfig.ClientSecret)
}
