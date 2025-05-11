package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AuthURL string `yaml:"auth_url"`
}

func LoadConfig(configPath string) (*Config, error) {

	if configPath == "" {
		return nil, fmt.Errorf("must provide a config path")
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	return &cfg, nil
}
