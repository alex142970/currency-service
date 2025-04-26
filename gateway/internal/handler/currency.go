package handler

import (
	"context"
	"currency_service/pkg/grpc_client"
	"github.com/gofiber/fiber/v2"
	"time"
)

type CurrencyRateResponse struct {
	Rate float32 `json:"rate"`
}

type CurrencyRateHistoryResponse struct {
	Rates []float32 `json:"rates"`
}

func CurrencyRateByDate(c *fiber.Ctx) error {
	inputDate := c.Query("date")

	date, _ := time.Parse("2006-01-02", inputDate)

	client, _ := grpc_client.NewCurrencyClient(context.Background(), "localhost:50051")

	rate, err := client.ExchangeRate(context.Background(), date)
	if err != nil {
		return err
	}

	return c.JSON(CurrencyRateResponse{
		Rate: rate.Rate,
	})
}

func CurrencyRateHistoryByDateInterval(c *fiber.Ctx) error {
	inputDateFrom := c.Query("dateFrom")
	inputDateTo := c.Query("dateTo")

	dateFrom, _ := time.Parse("2006-01-02", inputDateFrom)
	dateTo, _ := time.Parse("2006-01-02", inputDateTo)

	client, _ := grpc_client.NewCurrencyClient(context.Background(), "localhost:50051")

	rates, err := client.ExchangeRateHistory(context.Background(), dateFrom, dateTo)
	if err != nil {
		return err
	}

	return c.JSON(CurrencyRateHistoryResponse{
		Rates: rates.Rate,
	})
}
