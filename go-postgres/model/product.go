package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `gorm:"required" json:"title"`
	Description string `gorm:"required" json:"description"`
	Amount      int    `gorm:"required" json:"amount"`
}
