package handler

import (
	"fmt"

	"github.com/fahstjlps/sara-mongodb/database"
	"github.com/fahstjlps/sara-mongodb/model"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCollectionsLevel3(c *fiber.Ctx) error {
	// lookupStage := bson.D{{"$lookup", bson.D{{"from", "sub_sub_collections"}, {"localField", "_id"}, {"foreignField", "parent_id"}, {"as", "sub"}}}}

	// cursor, err := database.Mg.Db.Collection("sub_sub_collections").Aggregate(c.Context(), mongo.Pipeline{lookupStage})
	// if err != nil {
	// 	panic(err)
	// }

	// var showsLoaded []bson.M
	// if err = cursor.All(c.Context(), &showsLoaded); err != nil {
	// 	panic(err)
	// }

	// return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Query all collections success", "data": showsLoaded})
	return c.Status(200).JSON("xxx")
}

func CreateCollectionLevel3(c *fiber.Ctx) error {
	subCollection := database.Mg.Db.Collection("collections_level_2")
	subSubCollection := database.Mg.Db.Collection("collections_level_3")

	collectionLevel3 := new(model.CollectionLevel3)

	if err := c.BodyParser(collectionLevel3); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	insertionResult, err := subSubCollection.InsertOne(c.Context(), subSubCollection)
	fmt.Println(insertionResult)
	insertionResult2, err := subCollection.UpdateOne(c.Context(), bson.M{"_id": collectionLevel3.CollectionLevel2Id}, bson.D{{"$push", bson.D{{"collection_level_3", insertionResult.InsertedID}}}})
	fmt.Println(insertionResult2)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	id := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}

	createdRecord := subSubCollection.FindOne(c.Context(), id)

	createdSubSubCollection := &model.CollectionLevel3{}
	createdRecord.Decode(createdSubSubCollection)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Create a category success", "data": createdSubSubCollection})
}
