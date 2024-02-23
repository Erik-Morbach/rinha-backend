package config

type Config struct {
	DatabaseUrl string
}

func NewConfig() *Config {
	return &Config{
		DatabaseUrl: "postgresql://localhost:5432/rinha",
	}
}
