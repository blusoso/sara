package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryProduct struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CategoryId    primitive.ObjectID `json:"category_id,omitempty" bson:"category_id,omitempty"`
	SubCategoryId primitive.ObjectID `json:"sub_category_id,omitempty" bson:"sub_category_id,omitempty"`
	ProductId     primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	CreatedAt     time.Time          `json:"created_at" bson:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time          `json:"updated_at" bson:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
