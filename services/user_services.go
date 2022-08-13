package services

import (
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/auth"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/models"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/storage"
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
	if !valid {
		return context.Status(fiber.StatusUnauthorized).JSON(data)
	}

	user := models.User{
		FirstName: data["given_name"].(string),
		LastName:  data["family_name"].(string),
		Email:     data["email"].(string),
	}

	db := storage.GetDatabase()
	err = db.Create(&user).Error

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(user)
}
