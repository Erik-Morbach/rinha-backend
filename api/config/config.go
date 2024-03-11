package config

import "os"

type Config struct {
	DatabaseUrl string
	ServerPort string
}

func NewConfig() *Config {
	cfg := &Config{
		DatabaseUrl: "postgresql://localhost:5432/rinha",
		ServerPort: ":8080",
	}

	if len(os.Getenv("DB_URL")) > 0{
		cfg.DatabaseUrl = os.Getenv("DB_URL")
	}

	if len(os.Getenv("PORT")) > 0 {
		cfg.ServerPort = os.Getenv("PORT")
		if cfg.ServerPort[0] != ':' {
			cfg.ServerPort = ":" + cfg.ServerPort
		}
	}

	return cfg
}
