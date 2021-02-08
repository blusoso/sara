package handler

import (
	"github.com/fahstjlps/sara-mongodb/database"
	"github.com/fahstjlps/sara-mongodb/model"

	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
)

func GetCategories(c *fiber.Ctx) error {
	cursor, err := database.Mg.Db.Collection("categories").Find(c.Context(), bson.D{{}})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	var categories []model.Category
	if err := cursor.All(c.Context(), &categories); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Query all categories success", "data": categories})
}

func CreateCategory(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("categories")

	category := new(model.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	insertionResult, err := collection.InsertOne(c.Context(), category)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	id := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), id)

	createdCategory := &model.Category{}
	createdRecord.Decode(createdCategory)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Create a category success", "data": createdCategory})
}
