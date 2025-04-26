package handler

import (
	"currency_service/gateway/internal/clients/auth"
	"currency_service/gateway/internal/repository"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

func Login(c *fiber.Ctx) error {

	reqPayload := struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}{}

	err := c.BodyParser(&reqPayload)
	if err != nil {
		return err
	}

	usersRepo := repository.NewUserRepositoryMock()

	user := usersRepo.GetUserByLogin(reqPayload.Login)

	if user == nil {
		return c.JSON(fiber.Map{"error": "User not found"})
	}

	if user.Password != reqPayload.Password {
		return c.JSON(fiber.Map{"error": "Incorrect password"})
	}

	tokensClient := auth.NewTokensClient(&http.Client{
		Timeout: time.Second * 5,
	})

	token, err := tokensClient.Generate(user.Login)

	if err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}
