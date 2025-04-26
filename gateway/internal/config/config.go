package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env"`
}

func MustLoad(configPath string) *Config {

	if configPath == "" {
		panic("config file path not specified.")
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}
