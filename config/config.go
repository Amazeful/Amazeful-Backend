package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig *ServerConfig
	TwitchConfig *TwitchConfig
}

func LoadConfig() (*Config, error) {
	// load dotenv
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	config := &Config{
		ServerConfig: defaultServerConfig,
		TwitchConfig: defaultTwitchConfig,
	}

	err = env.Parse(config)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	err = validate.Struct(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
