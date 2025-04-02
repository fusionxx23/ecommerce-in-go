package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func ConnectDatabase() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Access environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	secretKey := os.Getenv("SECRET_KEY")
	apiKey := os.Getenv("API_KEY")

	// Use the variables
	fmt.Println("Database URL:", databaseURL)
	fmt.Println("Secret Key:", secretKey)
	fmt.Println("API Key:", apiKey)
	// Database connection string
	connStr := "postgres://username:password@localhost:5432/mydb"

	// Establishing the connection
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Successfully connected to the database!")

	// Running a simple query to test the connection
	var greeting string
	err = conn.QueryRow(context.Background(), "SELECT 'Hello, PostgreSQL!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	fmt.Println(greeting) // Output: Hello, PostgreSQL!
}
