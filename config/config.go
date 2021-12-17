package config

type ServerConfig struct {
	IpAddress string
	Port      string
	CertPath  string
	KeyPath   string
}

var initialConfig = &ServerConfig{
	IpAddress: "127.0.0.1",
	Port:      "8000",
	CertPath:  "",
	KeyPath:   "",
}

var config *ServerConfig = initialConfig

func GetConfig() *ServerConfig {
	return config
}
