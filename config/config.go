package config

import (
	"errors"
	"flag"
	"os"

	"github.com/joho/godotenv"
)

//ServerConfig includes all config data used by the web server
type ServerConfig struct {
	IpAddress string
	Port      string
	TLS       bool
	CertPath  string
	KeyPath   string
	MongoURI  string
}

var initialConfig = &ServerConfig{
	IpAddress: "127.0.0.1",
	Port:      "8000",
	CertPath:  "",
	KeyPath:   "",
}

var config *ServerConfig = initialConfig

//LoadConfig loads all config data into ServerConfig
func LoadConfig() error {
	flag.BoolVar(&config.TLS, "tls", false, "use ssl")
	flag.Parse()
	if config.TLS && (config.CertPath == "" || config.KeyPath == "") {
		return errors.New("to use ssl, you must provide CertPath and KeyPath")
	}
	err := godotenv.Load()
	if err != nil {
		return err
	}
	if os.Getenv("MONGO_URI") == "" {
		return errors.New("missing required env variable MONGO_URI")
	}
	config.MongoURI = os.Getenv("MONGO_URI")

	return nil
}

//GetConfig returns server config
func GetConfig() *ServerConfig {
	return config
}
