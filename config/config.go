package config

import (
	"flag"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

//ServerConfig includes all config data used by the web server
//and config coming from env variables
type ServerConfig struct {
	IpAddress     string `validate:"required,ip"`
	Port          string `validate:"required"`
	TLS           bool
	CertPath      string `validate:"required_with=TLS"`
	KeyPath       string `validate:"required_with=TLS"`
	Database      string `validate:"required"`
	MongoURI      string `validate:"required,uri"`
	RedisURI      string `validate:"required,uri"`
	RedisPassword string `validate:"required"`
	JwtSignKey    string `validate:"required"`
}

var initialConfig = &ServerConfig{
	IpAddress: "127.0.0.1",
	Port:      "8000",
	Database:  "Amazeful",
}

var config *ServerConfig = initialConfig

//LoadConfig loads all config data into ServerConfig
func LoadConfig() error {
	//Get the flags
	flag.BoolVar(&config.TLS, "tls", false, "use ssl")
	flag.Parse()

	// load dotenv
	err := godotenv.Load()
	if err != nil {
		return err
	}

	//set config
	config.MongoURI = os.Getenv("MONGO_URI")
	config.RedisURI = os.Getenv("REDIS_URI")
	config.RedisPassword = os.Getenv("REDIS_PASSWORD")
	config.JwtSignKey = os.Getenv("JWT_SIGN_KEY")
	config.CertPath = os.Getenv("CERT_PATH")
	config.KeyPath = os.Getenv("KEY_PATH")

	validate := validator.New()

	err = validate.Struct(config)
	if err != nil {
		return err
	}

	err = loadTwitchConfig()
	if err != nil {
		return err
	}

	return nil
}

//GetConfig returns server config
func GetConfig() *ServerConfig {
	return config
}
