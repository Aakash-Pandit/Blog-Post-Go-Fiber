package middleware

import (
	"fmt"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/auth"
	"github.com/gofiber/fiber/v2"
)

func GoogleAuthmiddleware() fiber.Handler {

	return func(context *fiber.Ctx) error {
		token := string(context.Request().Header.Peek("Authorization"))

		if token == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": "token is empty",
			})
		}

		valid, data := auth.GoogleTokenValidation(token, auth.GOOGLE_TOKEN_VALIDATION_URL)
		fmt.Println(data)
		if !valid {
			return context.Status(fiber.StatusUnauthorized).JSON(data)
		}

		return context.Next()
	}
}
