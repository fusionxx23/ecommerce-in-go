package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/fusionxx23/ecommerce-go/http/database"
	"github.com/fusionxx23/ecommerce-go/http/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Access environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	fmt.Println(databaseURL)
	// Establishing the connection

	database.DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Connected to database")
}

func SyncDb() {
	migrate := true
	if migrate {
		database.DB.AutoMigrate(&models.Cart{}, &models.CartItem{}, &models.Product{}, &models.User{}, &models.ProductVariant{}, &models.ProductImage{})
	}
	// change Chart id to text instead of int64 with GORM
	// database.DB.Migrator().ColumnTypes(&models.Cart{})
}
