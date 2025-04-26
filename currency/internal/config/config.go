package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string `yaml:"env"`
	Database DatabaseConfig
	Currency CurrencyConfig
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env:"DB_HOST" env-default:"localhost"`
	Port     int    `yaml:"port" env:"DB_PORT" env-default:"5432"`
	Database string `yaml:"database" env:"DB_DATABASE" env-default:"postgres"`
	User     string `yaml:"user" env:"DB_USER" env-default:"user"`
	Password string `yaml:"password" env:"DB_PASSWORD"`
}

type CurrencyConfig struct {
	Base   string `yaml:"base" env:"CURRENCY_BASE"`
	Target string `yaml:"target" env:"CURRENCY_TARGET"`
}

func (cfg *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
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
