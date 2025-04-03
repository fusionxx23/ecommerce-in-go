package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/fusionxx23/ecommerce-go/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	fmt.Println(databaseURL)
	// Establishing the connection
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}

func SyncDb() {
	DB.AutoMigrate(&models.Cart{}, &models.CartItem{}, &models.CheckoutSession{}, &models.DeliveryInfo{}, &models.Order{}, &models.Product{}, &models.User{})
}
