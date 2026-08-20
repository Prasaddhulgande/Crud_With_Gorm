package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Prasaddhulgande/Crud_With_Gorm/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// port=5433 is a docker port number
	//dsn := "host=127.0.0.1 port=5433 user=postgres password=postgres dbname=postgres sslmode=disable"
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Filed to connect to database:",err)
	}

	db.AutoMigrate(&models.User{})
	DB = db
}