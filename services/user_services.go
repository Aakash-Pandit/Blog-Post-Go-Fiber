package services

import (
	"fmt"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/auth"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/models"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/validators"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(context *fiber.Ctx) error {
	authtoken := models.AuthToken{}

	err := context.BodyParser(&authtoken)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	errors := validators.ValidateAuthTokenStruct(authtoken)

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	valid, data := auth.GoogleTokenValidation(authtoken.Token, auth.GOOGLE_TOKEN_VALIDATION_URL)
	fmt.Println(data)
	if !valid {
		return context.Status(fiber.StatusUnauthorized).JSON(data)
	}

	// user := models.User{}

	// err := context.BodyParser(&user)
	// if err != nil {
	// 	return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"detail": err.Error(),
	// 	})
	// }

	// errors := validators.ValidateUserStruct(user)

	// if errors != nil {
	// 	return context.Status(fiber.StatusBadRequest).JSON(errors)
	// }

	// repo := storage.GetDatabase()
	// err = repo.Create(&user).Error

	// if err != nil {
	// 	return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
	// 		"detail": err.Error(),
	// 	})
	// }

	// return context.Status(fiber.StatusCreated).JSON(user)

	return nil
}
