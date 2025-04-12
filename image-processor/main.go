package main

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
	"image"
	"image/jpeg"
	"log"

	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func init() {

	dsn := "host=localhost user=your_user password=your_password dbname=your_db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
func main() {
	fmt.Println("Consumer")
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
			if bounds.Dx() > bounds.Dy() {
				fmt.Println("The image is in landscape orientation.")
			} else {
				fmt.Println("The image is in portrait orientation.")
			}

			// Resize the image (example: 100x100)
			resizedImg400 := resize.Resize(420, 0, img, resize.Lanczos3)
			resizedImg130 := resize.Resize(130, 0, img, resize.Lanczos3)

			var buf bytes.Buffer
			err = jpeg.Encode(&buf, resizedImg400, nil)
			if err != nil {
				fmt.Printf("Failed to encode resized image: %v\n", err)
				continue
			}
			err = jpeg.Encode(&buf, resizedImg130, nil)
			if err != nil {
				fmt.Printf("Failed to encode resized image: %v\n", err)
				continue
			}
			fmt.Printf("Resized image size: %d bytes\n", buf.Len())
		}
	}()
	<-forever
}
