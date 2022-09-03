package validators

import (
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/models"
	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateBlogStruct(blog models.Blog) []*ErrorResponse {
	var errors []*ErrorResponse
	validation := validator.New()
	err := validation.Struct(blog)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateUserStruct(user models.User) []*ErrorResponse {
	var errors []*ErrorResponse
	validation := validator.New()
	err := validation.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateAuthTokenStruct(token models.AuthToken) []*ErrorResponse {
	var errors []*ErrorResponse
	validation := validator.New()
	err := validation.Struct(token)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func ValidateUsernameStruct(username models.Username) []*ErrorResponse {
	var errors []*ErrorResponse
	validation := validator.New()
	err := validation.Struct(username)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
