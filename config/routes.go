package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func HomePage(context *fiber.Ctx) error {
	var home = map[string]string{"BlogPost": "This is Home Page of BlogPost"}
	return context.Status(fiber.StatusOK).JSON(home)
}

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api/v1", logger.New())
	api.Get("/", HomePage)
}
