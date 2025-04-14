package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fusionxx23/ecommerce-go/image-processor/handlers"
	"github.com/fusionxx23/ecommerce-go/image-processor/libs"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var s3BucketName string

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()

	// Access environment variables
	s3BucketName = os.Getenv("S3_BUCKET_NAME")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Access environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	fmt.Println(databaseURL)
	// Establishing the connection

	libs.DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func main() {
	// get RABBITMQ_HOST from env
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	// check if rabbitmqHost is empty
	if rabbitmqHost == "" {
		rabbitmqHost = "localhost"
	}
	connUrl := fmt.Sprintf("amqp://guest:guest@%s:5672/", rabbitmqHost)
	conn, err := amqp.Dial(connUrl)

	if err != nil {
		fmt.Println(err, "ERROR")
		panic(err)
	}
	fmt.Println("Connected to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()
	msgs, err := ch.Consume("ImageQueue", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Println("recieved message")
			err := handlers.HandleCreateImage(d)
			if err != nil {
				continue
			}
		}
	}()
	<-forever
}

func uploadS3Image(b bytes.Buffer, key string) {

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession())

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3BucketName),
		Key:    aws.String(key),
		Body:   &b,
	})
	if err != nil {
		fmt.Printf("failed to upload file, %v", err)
		return
	}
	fmt.Printf("file uploaded to, %s\n", aws.StringValue(&result.Location))
}
