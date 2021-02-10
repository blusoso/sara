package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CollectionsProducts struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	CollectionLevel1Id primitive.ObjectID `json:"collection_level1_id,omitempty" bson:"collection_level1_id,omitempty"`
	CollectionLevel2Id primitive.ObjectID `json:"collection_level2_id" bson:"collection_level2_id"`
	CollectionLevel3Id primitive.ObjectID `json:"collection_level3_id" bson:"collection_level3_id"`
	ProductId          primitive.ObjectID `json:"product_id,omitempty" bson:"product_id,omitempty"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt          *time.Time         `json:"deleted_at" bson:"deleted_at"`
}
