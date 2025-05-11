package main

import (
	"currency_service/currency/internal/config"
	"errors"
	"flag"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

var (
	migrationsPath = flag.String("path", "", "Path to migrations folder")
	configPath     = flag.String("config", "config.yaml", "Path to config file")
)

func main() {
	flag.Parse()

	if *migrationsPath == "" {
		log.Fatalf("--path is required")
	}

	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal("failed to load config: %w", err)
	}

	m, err := migrate.New("file://"+*migrationsPath, conf.Database.GetDSN())
	if err != nil {
		log.Fatal("failed to connect to database: %w", err)
	}
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("failed to run migrations: %w", err)
	}
}
