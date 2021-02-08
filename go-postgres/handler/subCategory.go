package handler

import (
	"github.com/fahstjlps/go-sara/database"
	"github.com/fahstjlps/go-sara/model"

	"github.com/gofiber/fiber/v2"
)

func GetSubCategory(c *fiber.Ctx) error {
	db := database.DB
	var subCategories []model.SubCategory

	db.Preload("Category").Find(&subCategories)

	return c.JSON(fiber.Map{"status": "success", "message": "All sub categories", "data": subCategories})
}

func CreateSubCategory(c *fiber.Ctx) error {
	db := database.DB
	subCategory := new(model.SubCategory)

	if err := c.BodyParser(subCategory); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create sub category", "data": err})
	}

	db.Create(&subCategory)

	return c.JSON(fiber.Map{"status": "success", "message": "Create sub category", "data": subCategory})
}
