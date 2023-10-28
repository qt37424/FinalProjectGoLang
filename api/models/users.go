package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              uuid.UUID `gorm:"primary_key; unique;type:uuid; column:id"`
	Username        string    `gorm:"uniqueIndex"`
	Password        string
	Email           string
	TelephoneNumber string // Need prevent input is alphabet characters
	Avatar          string
	FirstName       string
	LastName        string
	IdentifyID      string // Maybe that's neccessary when renting car
	CarRent         []*Car `gorm:"-"` // List id car that user've rent
	Products        []Product
}
