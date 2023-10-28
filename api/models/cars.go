package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusEvents string

const (
	HIRING    StatusEvents = "IS_HIRING"    // This car is hiring by someone
	MAINTAIN  StatusEvents = "IS_MAINTAIN"  // it's maintaining you can't hire or doing anything
	AVAILABLE StatusEvents = "IS_AVAILABLE" // It's mean you can hire this car
)

type Car struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primary_key; unique;type:uuid; column:id"`
	NameCar     string
	Price       uint
	Status      StatusEvents
	UserRenting *User `gorm:"-"`
}
