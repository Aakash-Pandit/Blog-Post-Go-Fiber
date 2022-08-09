package routes

import (
	"github.com/Aakash-Pandit/Blog-Post-Go-Fiber/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/blogs", services.GetBlogs)
	api.Get("/blogs/:id", services.GetBlogByID)
	api.Post("/blogs", services.CreateBlog)
	api.Patch("/blogs/:id", services.UpdateBlog)
	api.Delete("/blogs/:id", services.DeleteBlog)
}
