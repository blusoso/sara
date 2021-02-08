package handler

import (
	"github.com/fahstjlps/sara-mongodb/database"
	"github.com/fahstjlps/sara-mongodb/model"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gofiber/fiber/v2"
)

func GetShops(c *fiber.Ctx) error {
	query := bson.D{{}}

	collection := database.Mg.Db.Collection("shops")
	cursor, err := collection.Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var shops []model.Shop
	if err := cursor.All(c.Context(), &shops); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(shops)
}

func CreateShop(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("shops")

	shop := new(model.Shop)

	if err := c.BodyParser(shop); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	insertionResult, err := collection.InsertOne(c.Context(), shop)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	id := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), id)

	createdShop := &model.Shop{}
	createdRecord.Decode(createdShop)

	return c.Status(201).JSON(createdShop)
}
