package handler

import (
	"github.com/fahstjlps/go-sara/database"
	"github.com/fahstjlps/go-sara/model"

	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {
	db := database.DB
	var products []model.Product

	titleQuery := c.Query("title")

	if titleQuery != "" {
		// db.Where("title = ?", titleQuery).Limit(2).Find(&products)
		db.Where("title = ?", titleQuery).Find(&products)

		if len(products) == 0 {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with title", "data": nil})
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Query the product success", "data": products})
	}

	db.Find(&products)

	return c.JSON(fiber.Map{"status": "success", "message": "All Products", "data": products})
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var product model.Product
	db.Find(&product, id)

	if product.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Product found", "data": product})
}

func CreateProduct(c *fiber.Ctx) error {
	db := database.DB
	product := new(model.Product) // &model.Product{} = Get the address of a struct

	if err := c.BodyParser(product); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}

	db.Create(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Create Product", "data": product})
}

func UpdateProduct(c *fiber.Ctx) error {
	type UpdateProductInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Amount      int    `json:"amount"`
	}

	newProduct := new(UpdateProductInput)

	if err := c.BodyParser(newProduct); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create product", "data": err})
	}

	db := database.DB
	id := c.Params("id")

	var product model.Product
	db.First(&product, id)

	if product.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	product.Title = newProduct.Title
	product.Description = newProduct.Description
	product.Amount = newProduct.Amount
	db.Save(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully updated", "data": product})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var product model.Product
	db.First(&product, id)

	if product.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}

	db.Delete(&product)

	return c.JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
