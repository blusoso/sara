package handler

import (
	"github.com/fahstjlps/sara-mongodb/database"
	"github.com/fahstjlps/sara-mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gofiber/fiber/v2"
)

func GetCollectionsProducts(c *fiber.Ctx) error {
	group := bson.D{{"$group", bson.D{{"_id", bson.D{{"collection_level1_id", "$collection_level1_id"}, {"collection_level2_id", "$collection_level2_id"}, {"collection_level3_id", "$collection_level3_id"}}}, {"products", bson.D{{"$push", "$product_id"}}}}}}
	project := bson.D{{"$project", bson.D{{"_id", 0}, {"collection_level1_id", "$_id.collection_level1_id"}, {"collection_level2_id", "$_id.collection_level2_id"}, {"collection_level3_id", "$_id.collection_level3_id"}, {"products", 1}}}}
	lookupCollection1 := bson.D{{"$lookup", bson.D{{"from", "collections_level_1"}, {"localField", "collection_level1_id"}, {"foreignField", "_id"}, {"as", "collection_level_1"}}}}
	lookupCollection2 := bson.D{{"$lookup", bson.D{{"from", "collections_level_2"}, {"localField", "collection_level2_id"}, {"foreignField", "_id"}, {"as", "collection_level_2"}}}}
	lookupCollection3 := bson.D{{"$lookup", bson.D{{"from", "collections_level_3"}, {"localField", "collection_level3_id"}, {"foreignField", "_id"}, {"as", "collection_level_3"}}}}

	showLoadedCursor, err := database.Mg.Db.Collection("collections_products").Aggregate(c.Context(), mongo.Pipeline{group, project, lookupCollection1, lookupCollection2, lookupCollection3})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	var showsLoaded []bson.M
	if err := showLoadedCursor.All(c.Context(), &showsLoaded); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Create a category success", "data": showsLoaded})
}

func CreateCollectionsProducts(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("collections_products")

	collectionsProducts := new(model.CollectionsProducts)
	if err := c.BodyParser(collectionsProducts); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	insertionResult, err := collection.InsertOne(c.Context(), collectionsProducts)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	id := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), id)

	createdCollectionsProducts := &model.CollectionsProducts{}
	createdRecord.Decode(createdCollectionsProducts)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Create a collections of products success", "data": createdCollectionsProducts})
}
