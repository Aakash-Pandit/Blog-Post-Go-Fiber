package middleware

import (
	"fmt"
	"strings"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/auth"
	"github.com/gofiber/fiber/v2"
)

func GoogleAuthmiddleware() fiber.Handler {

	return func(context *fiber.Ctx) error {
		token := string(context.Request().Header.Peek("Authorization"))
		token_info := strings.Split(token, " ")
		if len(token_info) != 2 {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": "length of the token should be 2",
			})
		}

		if token_info[0] != "Bearer" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": "Token should be Bearer",
			})
		}

		if token == "" {
			return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"detail": "token is empty",
			})
		}

		valid, data := auth.GoogleTokenValidation(token_info[1], auth.GOOGLE_TOKEN_VALIDATION_URL)
		fmt.Println(data)
		if !valid {
			return context.Status(fiber.StatusUnauthorized).JSON(data)
		}

		return context.Next()
	}
}
