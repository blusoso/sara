package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CollectionLevel2 struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name               string             `json:"name" bson:"name"`
	CollectionLevel1Id primitive.ObjectID `json:"collection_level_1_id" bson:"collection_level_1_id"`
	CollectionLevel3   []CollectionLevel3 `json:"collection_level_3" bson:"collection_level_3"`
	IsActive           bool               `json:"is_active" bson:"is_active"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt          *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
