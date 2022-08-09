package services

import (
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/models"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/storage"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/validators"
	"github.com/gofiber/fiber/v2"
)

func GetBlogs(context *fiber.Ctx) error {
	blogs := &[]models.Blog{}

	repo := storage.GetDatabase()
	err := repo.Find(blogs).Error
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

	repo := storage.GetDatabase()
	err := repo.Where("id = ?", id).First(blog).Error
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

	errors := validators.ValidateBlogStruct(blog)

	if errors != nil {
		return context.Status(fiber.StatusBadRequest).JSON(errors)
	}

	repo := storage.GetDatabase()
	err = repo.Create(&blog).Error

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

	repo := storage.GetDatabase()
	err = repo.Where("id = ?", id).Updates(blog).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"detail": err.Error(),
		})
	}

	err = repo.Where("id = ?", id).First(blog).Error
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

	repo := storage.GetDatabase()
	err := repo.Where("id = ?", id).First(blog).Error
	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"detail": err.Error(),
		})
	}

	repo.Where("id = ?", id).Delete(blog)

	return context.Status(fiber.StatusNoContent).JSON(&fiber.Map{})
}
