package main

import (
	application "currency_service/currency/internal/app"
	"currency_service/currency/internal/clients/currency"
	"currency_service/currency/internal/service"
	"fmt"
	"net/http"
	"time"
)

func main() {
	app := application.NewApp()

	client := currency_client.NewCurrencyClient(&http.Client{
		Timeout: time.Second * 5,
	})

	rate, err := client.FetchRate(app.Config.Currency.Base, app.Config.Currency.Target)
	if err != nil {
		panic(err)
	}

	s := service.NewCurrencyService(app)

	timeValue, _ := time.Parse("2006-01-02", rate.Date)

	err = s.AddRate(rate.Rate, timeValue)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
