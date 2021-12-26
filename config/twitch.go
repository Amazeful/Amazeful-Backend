package config

import (
	"os"

	"github.com/go-playground/validator/v10"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

type TwitchConfig struct {
	ClientID     string   `validate:"required"`
	ClientSecret string   `validate:"required"`
	Scopes       []string `validate:"required"`
	RedirectURL  string   `validate:"required"`
	State        string   `validate:"required"`
}

var initialTwitchConfig = &TwitchConfig{
	Scopes:      []string{"user:read:email"},
	RedirectURL: "http://localhost:8000/auth/twitch/callback",
}

var twitchConfig = initialTwitchConfig

//GetTwitchConfig returns twitch config
func GetTwitchConfig() *TwitchConfig {
	return twitchConfig
}

func GetTwitchOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     twitchConfig.ClientID,
		ClientSecret: twitchConfig.ClientSecret,
		Endpoint:     twitch.Endpoint,
		RedirectURL:  twitchConfig.RedirectURL,
		Scopes:       twitchConfig.Scopes,
	}
}

func loadTwitchConfig() error {
	twitchConfig.ClientID = os.Getenv("CLIENT_ID")
	twitchConfig.ClientSecret = os.Getenv("CLIENT_SECRET")
	twitchConfig.State = os.Getenv("STATE")

	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		return err
	}

	return nil
}
