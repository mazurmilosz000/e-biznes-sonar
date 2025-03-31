package database

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"api-project/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.Product{}, &models.Category{}, &models.Cart{})
}