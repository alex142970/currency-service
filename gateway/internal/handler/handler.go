package handler

import (
	"currency_service/gateway/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitHandlers(app *fiber.App) {
	app.Post("/auth/login", Login)

	app.Get("/currency/rate-history", middleware.AuthCheck, CurrencyRateHistoryByDateInterval)
	app.Get("/currency/rate", middleware.AuthCheck, CurrencyRateByDate)
}
