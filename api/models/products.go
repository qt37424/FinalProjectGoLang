package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID     uuid.UUID `gorm:"primary_key; unique;type:uuid; column:id"`
	Name   string
	Price  uint
	UserID uint
	User   User
}
