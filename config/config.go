package config

import "os"

type APIConfig struct {
	URL string
}

type Config struct {
	APIConfig
}

func (c *Config) readConfig() {
	api := os.Getenv("API_URL")
	c.APIConfig = APIConfig{
		URL: api,
	}
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}
