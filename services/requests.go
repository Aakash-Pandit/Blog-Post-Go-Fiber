package services

import (
	"errors"
	"strings"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/auth"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/models"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/storage"
	"github.com/gofiber/fiber/v2"
)

func FetchUserFromRequest(context *fiber.Ctx) (*models.User, error) {
	user := &models.User{}
	token := string(context.Request().Header.Peek("Authorization"))
	if token == "" {
		return &models.User{}, errors.New("token is empty")
	}

	token_info := strings.Split(token, " ")
	if len(token_info) != 2 {
		return &models.User{}, errors.New("length of the token should be 2")
	}

	if token_info[0] != "Bearer" {
		return &models.User{}, errors.New("token should be Bearer")
	}

	valid, data := auth.GoogleTokenValidation(token_info[1], auth.GOOGLE_TOKEN_VALIDATION_URL)
	if !valid {
		return &models.User{}, errors.New("invalid Token")
	}

	db := storage.GetDatabase()
	fetching_error := db.Where("email = ?", data["email"].(string)).First(user).Error
	if fetching_error != nil {
		return &models.User{}, errors.New("user does not exists")
	}

	return user, nil
}
