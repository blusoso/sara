package handler

import (
	"github.com/fahstjlps/go-sara/database"
	"github.com/fahstjlps/go-sara/model"

	"github.com/gofiber/fiber/v2"
)

func GetCategory(c *fiber.Ctx) error {
	db := database.DB
	var categories []model.Category

	db.Find(&categories)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "All Categories", "data": categories})
}

func CreateCategory(c *fiber.Ctx) error {
	db := database.DB
	category := new(model.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create category", "data": err})
	}

	db.Create(&category)

	return c.JSON(fiber.Map{"status": "success", "message": "Create category", "data": category})
}
