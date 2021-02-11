package handler

import (
	"time"

	"github.com/fahstjlps/sara-mongodb/database"
	"github.com/fahstjlps/sara-mongodb/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gofiber/fiber/v2"
)

func GetVenderCompanies(c *fiber.Ctx) error {
	opts := options.Find()
	opts.SetSort(bson.D{{"_id", -1}})

	collection := database.Mg.Db.Collection("vender-companies")
	query := bson.D{{"deleted_at", bson.D{{"$eq", nil}}}}

	cursor, err := collection.Find(c.Context(), query, opts)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	var venderCompanies []model.VenderCompany
	if err := cursor.All(c.Context(), &venderCompanies); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Query all vender companies success", "data": venderCompanies})
}

func GetVenderCompany(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("vender-companies")
	idParam := c.Params("id")
	VenderCompanyId, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	id := bson.M{"_id": VenderCompanyId}
	queryRecord := collection.FindOne(c.Context(), id)

	venderCompany := &model.VenderCompany{}
	queryRecord.Decode(venderCompany)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Query a vender company success", "data": venderCompany})
}

func CreateVenderCompany(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("vender-companies")
	venderCompany := new(model.VenderCompany)

	if err := c.BodyParser(venderCompany); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	insertionResult, err := collection.InsertOne(c.Context(), venderCompany)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	id := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), id)

	createdVenderCompany := &model.VenderCompany{}
	createdRecord.Decode(createdVenderCompany)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Vender companies successfully created", "data": createdVenderCompany})
}

func UpdateVenderCompany(c *fiber.Ctx) error {
	idParam := c.Params("id")
	venderCompanyId, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	venderCompany := new(model.VenderCompany)

	if err := c.BodyParser(venderCompany); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	id := bson.D{{Key: "_id", Value: venderCompanyId}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: venderCompany.Name},
				{Key: "address", Value: venderCompany.Address},
				{Key: "phone_number", Value: venderCompany.PhoneNumber},
			},
		},
	}

	err = database.Mg.Db.Collection("vender-companies").FindOneAndUpdate(c.Context(), id, update).Err()
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Vender company successfully updated", "data": venderCompany})
}

func DeleteVenderCompany(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("vender-companies")
	venderCompanyID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	id := bson.D{{Key: "_id", Value: venderCompanyID}}
	updateDeletedAt := bson.D{{Key: "$set", Value: bson.D{{Key: "deleted_at", Value: time.Now()}}}}

	err = collection.FindOneAndUpdate(c.Context(), id, updateDeletedAt).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Vender company successfully deleted", "data": nil})
}
