package handler

import (
	"github.com/fahstjlps/sara-mongodb/database"
	"github.com/fahstjlps/sara-mongodb/model"

	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCollectionsLevel2(c *fiber.Ctx) error {
	lookupStage := bson.D{{"$lookup", bson.D{{"from", "collections_level_2"}, {"localField", "_id"}, {"foreignField", "collection_level_2_id"}, {"as", "sub_collection"}}}}

	cursor, err := database.Mg.Db.Collection("collections_level_2").Aggregate(c.Context(), mongo.Pipeline{lookupStage})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	var showsLoaded []bson.M
	if err = cursor.All(c.Context(), &showsLoaded); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Query all collections success", "data": showsLoaded})
}

func CreateCollectionLevel2(c *fiber.Ctx) error {
	database := database.Mg.Db.Collection("collections_level_2")

	collection := new(model.CollectionLevel2)

	if err := c.BodyParser(collection); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	insertionResult, err := database.InsertOne(c.Context(), collection)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	id := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

	createdRecord := database.FindOne(c.Context(), id)

	createCollection := &model.CollectionLevel2{}
	createdRecord.Decode(createCollection)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Create a category success", "data": createCollection})
}
