package handler

import (
	"strconv"
	"time"

	"github.com/fahstjlps/sara-mongodb/database"
	"github.com/fahstjlps/sara-mongodb/model"

	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetProducts(c *fiber.Ctx) error {
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	limit := int64(10)
	var offset int64

	if page > 0 {
		offset = (page - 1) * limit
	}

	opts := options.Find()
	opts.SetSort(bson.D{{"_id", -1}})
	opts.SetLimit(limit)
	opts.SetSkip(offset)

	collection := database.Mg.Db.Collection("products")
	query := bson.D{{"deleted_at", bson.D{{"$eq", nil}}}}
	cursor, err := collection.Find(c.Context(), query, opts)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	var products []model.Product

	if err := cursor.All(c.Context(), &products); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Query all products success", "data": products})
}

func GetProduct(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("products")
	idParam := c.Params("id")
	productID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	id := bson.M{"_id": productID}
	queryRecord := collection.FindOne(c.Context(), id)

	product := &model.Product{}
	queryRecord.Decode(product)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Query a product success", "data": product})
}

func CreateProduct(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("products")

	// New Product struct
	product := new(model.Product)
	// Parse body into struct
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	// insert the record
	insertionResult, err := collection.InsertOne(c.Context(), product)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	// get the just inserted record in order to return it as response
	id := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), id)

	createdProduct := &model.Product{}
	createdRecord.Decode(createdProduct)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product successfully created", "data": createdProduct})
}

func UpdateProduct(c *fiber.Ctx) error {
	idParam := c.Params("id")
	productID, err := primitive.ObjectIDFromHex(idParam)

	// the provided ID might be invalid ObjectID
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	product := new(model.Product)
	// Parse body into struct
	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	// Find the product and update its data
	id := bson.D{{Key: "_id", Value: productID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "sku", Value: product.Sku},
				{Key: "title", Value: product.Title},
				{Key: "brand", Value: product.Brand},
				{Key: "condition", Value: product.Condition},
				{Key: "meta_description", Value: product.MetaDescription},
				{Key: "description", Value: product.Description},
				{Key: "colors", Value: product.Colors},
				{Key: "sizes", Value: product.Sizes},
				{Key: "price", Value: product.Price},
				{Key: "qty", Value: product.Qty},
				{Key: "warranty_id", Value: product.WarrantyId},
				{Key: "shop", Value: product.Shop},
				{Key: "is_active", Value: product.IsActive},
				{Key: "publish_date", Value: product.PublishDate},
			},
		},
	}
	err = database.Mg.Db.Collection("products").FindOneAndUpdate(c.Context(), id, update).Err()

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product successfully updated", "data": product})
}

func DeleteProduct(c *fiber.Ctx) error {
	collection := database.Mg.Db.Collection("products")
	productID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	id := bson.D{{Key: "_id", Value: productID}}
	updateDeletedAt := bson.D{{Key: "$set", Value: bson.D{{Key: "deleted_at", Value: time.Now()}}}}

	err = collection.FindOneAndUpdate(c.Context(), id, updateDeletedAt).Err()

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
		}
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error(), "data": ""})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Product successfully deleted", "data": nil})
}
