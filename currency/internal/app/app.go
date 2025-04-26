package app

import (
	"currency_service/currency/internal/config"
	"currency_service/currency/internal/db"
	"database/sql"
	"flag"
)

type App struct {
	Config *config.Config
	Db     *sql.DB
}

var (
	configPath = flag.String("config", "", "path to the configuration file")
)

func NewApp() *App {
	flag.Parse()

	conf := config.MustLoad(*configPath)
	connection := db.Connect(&conf.Database)

	return &App{
		Config: conf,
		Db:     connection,
	}
}
