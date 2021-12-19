package config

import (
	"errors"
	"flag"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
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
	Twitch      *Twitch
}

type Twitch struct {
	ClientID     string
	ClientSecret string
	Scopes       []string
	RedirectURL  string
	State        string
	OauthConfig  *oauth2.Config
}

var initialConfig = &ServerConfig{
	IpAddress: "127.0.0.1",
	Port:      "8000",
	CertPath:  "",
	KeyPath:   "",
	Twitch: &Twitch{
		Scopes:      []string{"user:read:email"},
		RedirectURL: "http://localhost:8000/auth/twitch/callback",
	},
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
	config.Twitch.ClientID = os.Getenv("CLIENT_ID")
	config.Twitch.ClientSecret = os.Getenv("CLIENT_SECRET")
	config.Twitch.State = os.Getenv("STATE")
	config.TokenSecret = os.Getenv("TOKEN_SECRET")

	config.Twitch.OauthConfig = &oauth2.Config{
		ClientID:     config.Twitch.ClientID,
		ClientSecret: config.Twitch.ClientSecret,
		Endpoint:     twitch.Endpoint,
		RedirectURL:  config.Twitch.RedirectURL,
		Scopes:       config.Twitch.Scopes,
	}

	return nil
}

//GetConfig returns server config
func GetConfig() *ServerConfig {
	return config
}

//GetTwitchConfig returns twitch config
func GetTwitchConfig() *Twitch {
	return config.Twitch
}
