package main

import (
	"FinalProject/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	// Connect to PostgreDB
	dsn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// migrate data into database (that is neccessary to connect database)
	db.AutoMigrate(&models.User{}, &models.Product{}, &models.Car{})

	return db
}
