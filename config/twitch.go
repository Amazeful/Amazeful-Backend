package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

type TwitchConfig struct {
	ClientID     string `env:"TWITCH_CLIENT_ID" validate:"required"`
	ClientSecret string `env:"TWITCH_CLIENT_SECRET" validate:"required"`
	State        string `env:"TWITCH_STATE" validate:"required"`
}

func (c *Config) GetTwitchOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     c.TwitchConfig.ClientID,
		ClientSecret: c.TwitchConfig.ClientSecret,
		Endpoint:     twitch.Endpoint,
		RedirectURL:  c.ServerConfig.ServerURL + "/auth/twitch/callback",
		Scopes:       []string{"user:read:email"},
	}
}

var defaultTwitchConfig = &TwitchConfig{
	State: "test_state",
}
