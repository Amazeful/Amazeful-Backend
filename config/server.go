package config

type ServerConfig struct {
	IpAddress     string `env:"IP" validate:"required,ip"`
	ServerURL     string `env:"SERVER_URL" validate:"required,url"`
	Port          string `env:"SERVER_PORT" validate:"required"`
	TLS           bool   `env:"TLS"`
	CertPath      string `env:"CERT_PATH" validate:"required_with=TLS"`
	KeyPath       string `env:"KEY_PATH" validate:"required_with=TLS"`
	MongoURI      string `env:"MONGO_URI" validate:"required,uri"`
	RedisURI      string `env:"REDIS_URI" validate:"required,uri"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	JwtSignKey    string `env:"JWT_SIGN_KEY" validate:"required"`
}

var defaultServerConfig = &ServerConfig{
	IpAddress:  "127.0.0.1",
	ServerURL:  "http://localhost:8000",
	MongoURI:   "localhost:27017",
	RedisURI:   "localhost:6379",
	Port:       "8000",
	JwtSignKey: "test_key",
}
