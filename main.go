package main

import (
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/config"
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	setupConfig := config.GetApplicationConfig()
	database.InitializeDB(setupConfig)
	db := database.GetDatabase()

	config.SetupRoutes(app, db)

	app.Listen(":" + setupConfig.BackendPort)
}
