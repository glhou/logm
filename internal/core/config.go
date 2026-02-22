package core

import "github.com/caarlos0/env/v11"

type Config struct {
	DatabaseUrl string `env:"DATABASE_URL",required`
}

func loadConfig() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
