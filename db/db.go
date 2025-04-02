package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Access environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	// Establishing the connection
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

}
