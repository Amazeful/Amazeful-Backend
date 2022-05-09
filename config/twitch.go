package config

import (
	"github.com/caarlos0/env/v6"
	validator "github.com/go-playground/validator/v10"
)

type TwitchConfig struct {
	ClientID     string `env:"TWITCH_CLIENT_ID" validate:"required"`
	ClientSecret string `env:"TWITCH_CLIENT_SECRET" validate:"required"`
	State        string `env:"TWITCH_STATE" validate:"required"`
}

func NewTwitchConfig(serverUrl string) *TwitchConfig {
	return &TwitchConfig{
		State: "test_state",
	}
}

func (tc *TwitchConfig) Load() error {
	err := env.Parse(tc)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(tc)
	if err != nil {
		return err
	}

	return nil
}
