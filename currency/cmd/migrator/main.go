package main

import (
	application "currency_service/currency/internal/app"
	"flag"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

var (
	migrationsPath = flag.String("path", "", "Path to migrations folder")
)

func main() {

	app := application.NewApp()

	if *migrationsPath == "" {
		panic("migrations folder not specified.")
	}

	m, err := migrate.New("file://"+*migrationsPath, app.Config.Database.GetDSN())
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}
