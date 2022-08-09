package main

import (
	"log"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/config"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/models"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/routes"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	config := config.SetupEnv()
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not load database")
	}

	err = models.MigrateBlogs(db)
	if err != nil {
		log.Fatal("could not migrate db")
	}

	app := fiber.New()
	app.Use(logger.New())
	routes.SetupRoutes(app)

	app.Listen(":" + config.BackendPort)
}
