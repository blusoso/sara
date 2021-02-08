package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID              primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Sku             string             `json:"sku" bson:"sku"`
	Title           string             `json:"title" bson:"title"`
	Brand           string             `json:"brand" bson:"brand"`
	Condition       string             `json:"condition" bson:"condition"`
	MetaDescription string             `json:"meta_description" bson:"meta_description"`
	Description     string             `json:"description" bson:"description"`
	Colors          []string           `json:"colors" bson:"colors"`
	Sizes           []string           `json:"sizes" bson:"sizes"`
	Price           float32            `json:"price" bson:"price"`
	Qty             uint               `json:"qty" bson:"qty"`
	WarrantyId      int                `json:"warranty_id" bson:"warranty_id"`
	Shop            primitive.ObjectID `json:"shop,omitempty" bson:"shop"`
	Category        primitive.ObjectID `json:"category" bson:"category"`
	IsActive        bool               `json:"is_active" bson:"is_active"`
	PublishDate     time.Time          `json:"publish_date" bson:"publish_date"`
	CreatedAt       time.Time          `json:"created_at" bson:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt       time.Time          `json:"updated_at" bson:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt       *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
