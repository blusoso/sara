package router

import (
	"fmt"

	"github.com/fahstjlps/go-sara/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", func(c *fiber.Ctx) error {
		fmt.Println("🥇 API Middleware")
		return c.Next()
	})

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		fmt.Println("🥇 API v.1 Middleware")
		return c.Next()
	})

	// Products
	product := v1.Group("/product")
	product.Get("/", handler.GetProducts)
	product.Get("/:id", handler.GetProduct)
	product.Post("/", handler.CreateProduct)
	product.Put("/:id", handler.UpdateProduct)
	product.Delete("/:id", handler.DeleteProduct)

	// Categories
	category := v1.Group("/category")
	category.Get("/", handler.GetCategory)
	category.Post("/", handler.CreateCategory)

	// Sub Categories
	subCategory := v1.Group("/sub-category")
	subCategory.Get("/", handler.GetSubCategory)
	subCategory.Post("/", handler.CreateSubCategory)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
}
