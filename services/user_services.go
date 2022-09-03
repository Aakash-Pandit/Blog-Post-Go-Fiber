package services

import (
	"fmt"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/auth"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/models"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/storage"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/validators"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(context *fiber.Ctx) error {
	users := &[]models.User{}

	db := storage.GetDatabase()
	err := db.Find(users).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"count":   len(*users),
		"results": users,
	})
}

func GetUserByID(context *fiber.Ctx) error {
	id := context.Params("id")
	user := &models.User{}

	db := storage.GetDatabase()
	err := db.Where("id = ?", id).First(user).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(user)
}

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

	db := storage.GetDatabase()
	user := &models.User{}
	fetching_error := db.Where("email = ?", data["email"].(string)).First(user).Error
	if fetching_error != nil {
		new_user := models.User{
			FirstName: data["given_name"].(string),
			LastName:  data["family_name"].(string),
			Email:     data["email"].(string),
		}

		err = db.Create(&new_user).Error
		if err != nil {
			return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
				"detail": err.Error(),
			})
		}

		return context.Status(fiber.StatusCreated).JSON(new_user)
	}

	return context.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	username := models.Username{}

	err := context.BodyParser(&username)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	errors := validators.ValidateUsernameStruct(username)

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db := storage.GetDatabase()
	user := &models.User{}
	err = db.Where("username = ?", username.Username).First(user).Error
	if err == nil {
		fmt.Println("err: ", err)
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"detail": "Username already exists.",
		})
	}

	user, err = FetchUserFromRequest(context)
	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(err)
	}

	if id != user.ID.String() {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": "you are not authorized",
		})
	}

	user.Username = username.Username
	err = db.Where("id = ?", id).Updates(user).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(user)
}
