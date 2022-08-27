package services

import (
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/models"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/storage"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/validators"
	"github.com/gofiber/fiber/v2"
)

func GetBlogs(context *fiber.Ctx) error {
	blogs := &[]models.Blog{}

	db := storage.GetDatabase()
	err := db.Find(blogs).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(&fiber.Map{
		"count":   len(*blogs),
		"results": blogs,
	})
}

func GetBlogByID(context *fiber.Ctx) error {
	id := context.Params("id")
	blog := &models.Blog{}

	db := storage.GetDatabase()
	err := db.Where("id = ?", id).First(blog).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(blog)
}

func CreateBlog(context *fiber.Ctx) error {
	blog := models.Blog{}

	err := context.BodyParser(&blog)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	user, err := FetchUserFromRequest(context)
	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(err)
	}

	blog.CreatedById = user.ID

	errors := validators.ValidateBlogStruct(blog)

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	db := storage.GetDatabase()
	err = db.Create(&blog).Error

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusCreated).JSON(blog)
}

func UpdateBlog(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	blog := &models.Blog{}

	err := context.BodyParser(&blog)
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	errors := validators.ValidateBlogStruct(*blog)
	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	user, err := FetchUserFromRequest(context)
	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(err)
	}

	if blog.CreatedById != user.ID {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": "you are not authorized",
		})
	}

	db := storage.GetDatabase()
	err = db.Where("id = ?", id).Updates(blog).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"detail": err.Error(),
		})
	}

	err = db.Where("id = ?", id).First(blog).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	return context.Status(fiber.StatusOK).JSON(blog)
}

func DeleteBlog(context *fiber.Ctx) error {
	blog := &models.Blog{}
	id := context.Params("id")

	db := storage.GetDatabase()
	err := db.Where("id = ?", id).First(blog).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	user, err := FetchUserFromRequest(context)
	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(err)
	}

	if blog.CreatedById != user.ID {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"detail": "you are not authorized",
		})
	}

	db.Where("id = ?", id).Delete(blog)

	return context.Status(fiber.StatusNoContent).JSON(&fiber.Map{})
}
