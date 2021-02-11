package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VenderCompany struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Address     string             `json:"address" bson:"address"`
	PhoneNumber string             `json:"phone_number" bson:"phone_number"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
