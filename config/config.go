package config

import (
	"fmt"
	"log"
	"os"

	"go-jwt-auth/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the database URL from environment variable
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// Open the database connection using GORM (GORM v2)
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	fmt.Println("Database connected successfully")

	// Auto-migrate the database (create tables if they don't exist)
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error during AutoMigrate:", err)
	}
}

// CloseDatabase closes the database connection
func CloseDatabase() {
	sqlDB, err := DB.DB() 
	if err != nil {
		log.Fatal("Error getting database connection:", err)
	}
	sqlDB.Close() 
}
