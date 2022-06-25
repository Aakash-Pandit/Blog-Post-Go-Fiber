package main

import (
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})
	config.ReadEnv()
	app.Listen(":8080")
}
