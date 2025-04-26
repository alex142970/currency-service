package main

import (
	"currency_service/gateway/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	handler.InitHandlers(app)

	app.Listen(":3000")
}
