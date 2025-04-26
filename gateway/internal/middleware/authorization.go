package middleware

import (
	"currency_service/gateway/internal/clients/auth"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strings"
	"time"
)

func AuthCheck(ctx *fiber.Ctx) error {
	bearerToken := ctx.Get("Authorization")

	token := strings.TrimSpace(strings.Replace(bearerToken, "Bearer", "", 1))

	tokensClient := auth.NewTokensClient(&http.Client{Timeout: 5 * time.Second})

	isValid := tokensClient.Validate(token)

	fmt.Println(isValid)

	if !isValid {
		ctx.Status(fiber.StatusUnauthorized)
		ctx.JSON(fiber.Map{"error": "Invalid token"})
		return nil
	}

	return ctx.Next()
}
