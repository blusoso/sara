package model

import (
	"gorm.io/gorm"
)

type SubCategory struct {
	gorm.Model
	Name       string   `gorm:"required" json:"name"`
	CategoryID uint     `gorm:"required" json:"category_id"`
	Category   Category `json:"category"`
}
