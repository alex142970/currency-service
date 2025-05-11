package main

import (
	"context"
	"currency_service/currency/internal/clients/currency"
	"currency_service/currency/internal/config"
	"currency_service/currency/internal/db"
	"currency_service/currency/internal/repository"
	"currency_service/currency/internal/service"
	"flag"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var (
	configPath = flag.String("config", "config.yaml", "Path to config file")
)

func main() {
	flag.Parse()

	conf, err := config.LoadConfig(*configPath)

	if err != nil {
		log.Fatal("Error loading config: %w", err)
	}

	client := currency_client.NewCurrencyClient(&http.Client{
		Timeout: time.Second * 5,
	})

	rate, err := client.FetchRate(conf.Currency.Base, conf.Currency.Target)
	if err != nil {
		log.Fatal("Error fetching rate: %w", err)
	}

	connection, err := db.NewDatabaseConnection(&conf.Database)

	repo := repository.NewCurrencyPostgresRepository(connection)

	logger := slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelDebug)

	s := service.NewCurrencyService(repo, conf, logger)

	timeValue, _ := time.Parse("2006-01-02", rate.Date)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = s.AddRate(ctx, rate.Rate, timeValue)
	if err != nil {
		log.Fatal("Error adding rate: %w", err)
	}
}
