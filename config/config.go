package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
)

var config *Config

type Config struct {
	ServerConfig *ServerConfig
	TwitchConfig *TwitchConfig
}

func LoadConfig() error {
	cfg := &Config{
		ServerConfig: defaultServerConfig,
		TwitchConfig: defaultTwitchConfig,
	}

	err := env.Parse(cfg)
	if err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return err
	}

	config = cfg

	return nil
}

func GetConfig() *Config {
	return config
}
