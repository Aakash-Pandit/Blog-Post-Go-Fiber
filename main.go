package main

import (
	"log"

	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/config"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/middleware"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/routes"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	configuration := config.SetupEnv()
	db, err := storage.NewConnection(configuration)
	if err != nil {
		log.Fatal("could not load database")
	}

	err = storage.MigrateDatabase(db)
	if err != nil {
		return
	}

	app := fiber.New()
	app.Use("", logger.New())
	app.Use("/api/v1", logger.New(), middleware.GoogleAuthmiddleware())
	routes.SetupRoutes(app)

	app.Listen(":" + configuration.BackendPort)
}
