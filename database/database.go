package database

import (
	"CryptogameBackTs/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	fmt.Println("Database connection successfully opened")

	// Миграция моделей
	err = DB.AutoMigrate(&models.Author{}, &models.Level{}, &models.Quote{}, &models.User{}, &models.AccessLevel{}, &models.Session{})

	if err != nil {
		log.Fatal("failed to migrate database")
	}

	fmt.Println("Database migration completed")
}
