package router

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/fahstjlps/sara-mongodb/handler"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())

	api := app.Group("/api", func(c *fiber.Ctx) error {
		fmt.Println("🥇 API Middleware")
		return c.Next()
	})

	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		fmt.Println("🥇 API v.1 Middleware")
		return c.Next()
	})

	product := v1.Group("/product")
	product.Get("/", handler.GetProducts)
	product.Get("/:id", handler.GetProduct)
	product.Post("/", handler.CreateProduct)
	product.Put("/:id", handler.UpdateProduct)
	product.Delete("/:id", handler.DeleteProduct)

	shop := v1.Group("/shop")
	shop.Get("/", handler.GetShops)
	shop.Post("/", handler.CreateShop)

	collectionLevel1 := v1.Group("/collection-level-1")
	collectionLevel1.Get("/", handler.GetCollectionsLevel1)
	collectionLevel1.Post("/", handler.CreateCollectionLevel1)

	collectionLevel2 := v1.Group("/collection-level-2")
	collectionLevel2.Get("/", handler.GetCollectionsLevel2)
	collectionLevel2.Post("/", handler.CreateCollectionLevel2)

	collectionLevel3 := v1.Group("/collection-level-3")
	collectionLevel3.Get("/", handler.GetCollectionsLevel3)
	collectionLevel3.Post("/", handler.CreateCollectionLevel3)

	category := v1.Group("/category")
	category.Get("/", handler.GetCategories)
	category.Post("/", handler.CreateCategory)
}
