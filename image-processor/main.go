package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/fusionxx23/ecommerce-go/image-processor/models"
	"github.com/joho/godotenv"
	"github.com/nfnt/resize"
	"github.com/streadway/amqp"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB
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

	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()
	msgs, err := ch.Consume("TestQueue", "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {

			headers := d.Headers
			imageID, ok := headers["image-id"]
			if !ok {
				log.Println("image-id header not found")
				continue
			}
			log.Printf("Image ID: %v", imageID)
			img, _, err := image.Decode(bytes.NewReader(d.Body))
			if err != nil {
				fmt.Println("Error decoding image:", err)
				continue
			}

			// Determine if the image is portrait or landscape
			bounds := img.Bounds()
			orientation := ""
			if bounds.Dx() > bounds.Dy() {
				fmt.Println("The image is in landscape orientation.")
				orientation = "landscape"
			} else {
				orientation = "portrait"
				fmt.Println("The image is in portrait orientation.")
			}

			if orientation == "landscape" { // no need to optimize landscape picture
				models.UpdateProductImage(DB, imageID.(string), orientation)
				continue
			}

			// Resize the image (example: 100x100)
			resizedImg1260 := resize.Resize(1260, 0, img, resize.Lanczos3)
			resizedImg420 := resize.Resize(420, 0, img, resize.Lanczos3)
			resizedImg130 := resize.Resize(130, 0, img, resize.Lanczos3)

			var img1260Buffer bytes.Buffer
			var img420Buffer bytes.Buffer
			var img130Buffer bytes.Buffer
			err = jpeg.Encode(&img420Buffer, resizedImg420, nil)
			if err != nil {
				fmt.Printf("Failed to encode resized image420: %v\n", err)
				continue
			}
			err = jpeg.Encode(&img130Buffer, resizedImg130, nil)
			if err != nil {
				fmt.Printf("Failed to encode resized image130: %v\n", err)
				continue
			}
			err = jpeg.Encode(&img1260Buffer, resizedImg1260, nil)
			if err != nil {
				fmt.Printf("Failed to encode resized image 1260: %v\n", err)
				continue
			}
			uploadS3Image(img130Buffer, imageID.(string)+"-130.jpg")
			uploadS3Image(img420Buffer, imageID.(string)+"-420.jpg")
			uploadS3Image(img1260Buffer, imageID.(string)+"-1260.jpg")
			models.UpdateProductImage(DB, imageID.(string), orientation)
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
