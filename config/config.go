package config

import (
	"errors"
	"flag"
	"os"

	"github.com/joho/godotenv"
)

//ServerConfig includes all config data used by the web server
//and config coming from env variables
type ServerConfig struct {
	IpAddress   string
	Port        string
	TLS         bool
	CertPath    string
	KeyPath     string
	MongoURI    string
	TokenSecret string
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
	//Get the flags
	flag.BoolVar(&config.TLS, "tls", false, "use ssl")
	flag.Parse()

	//check cert
	if config.TLS && (config.CertPath == "" || config.KeyPath == "") {
		return errors.New("to use ssl, you must provide CertPath and KeyPath")
	}

	// load dotenv
	err := godotenv.Load()
	if err != nil {
		return err
	}

	//check db info
	if os.Getenv("MONGO_URI") == "" {
		return errors.New("missing required env variable MONGO_URI")
	}

	//set config
	config.MongoURI = os.Getenv("MONGO_URI")
	config.TokenSecret = os.Getenv("TOKEN_SECRET")

	loadTwitchConfig()

	return nil
}

//GetConfig returns server config
func GetConfig() *ServerConfig {
	return config
}
