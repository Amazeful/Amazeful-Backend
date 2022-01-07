package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var config *Config

type Config struct {
	ServerConfig *ServerConfig
	TwitchConfig *TwitchConfig
}

func LoadConfig() error {
	// load dotenv
	err := godotenv.Load()
	if err != nil {
		return err
	}
	cfg := &Config{
		ServerConfig: defaultServerConfig,
		TwitchConfig: defaultTwitchConfig,
	}

	err = env.Parse(cfg)
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
