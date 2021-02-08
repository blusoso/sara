package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Shop struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	PhoneNumber string             `json:"phone_number"`
	CreatedAt   time.Time          `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time          `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time         `json:"deleted_at"`
}
