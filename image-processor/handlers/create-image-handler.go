package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/jpeg"
	"log"

	"golang.org/x/image/webp"

	"github.com/fusionxx23/ecommerce-go/image-processor/libs"
	"github.com/fusionxx23/ecommerce-go/image-processor/models"
	"github.com/nfnt/resize"
	"github.com/streadway/amqp"
)

type ImagePayload struct {
	Name  string `json:"name"`
	Bytes string `json:"bytes"`
}

func HandleCreateImage(d amqp.Delivery) error {

	var payload ImagePayload
	err := json.Unmarshal(d.Body, &payload)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return fmt.Errorf("error parsing JSON")
	}
	fmt.Println("Unmarshalled")
	imageBytes, err := decodeBase64(payload.Bytes)
	if err != nil {
		fmt.Println("Error decoding base 64:", err)
		return fmt.Errorf("decoding base 64")
	}
	img, err := webp.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return fmt.Errorf("decoding image")
	}

	fmt.Println("Decoded")
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
		models.UpdateProductImage(libs.DB, payload.Name, orientation)

		return fmt.Errorf("cannot divide by zero")
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
		return fmt.Errorf("failed to encode")
	}
	err = jpeg.Encode(&img130Buffer, resizedImg130, nil)
	if err != nil {
		fmt.Printf("Failed to encode resized image130: %v\n", err)
		return fmt.Errorf("failed to encode")
	}
	err = jpeg.Encode(&img1260Buffer, resizedImg1260, nil)
	if err != nil {
		fmt.Printf("Failed to encode resized image 1260: %v\n", err)
		return fmt.Errorf("failed to encode")
	}
	libs.UploadS3Image(img130Buffer, payload.Name+"-130.webp")
	libs.UploadS3Image(img420Buffer, payload.Name+"-420.webp")
	libs.UploadS3Image(img1260Buffer, payload.Name+"-1260.webp")
	models.UpdateProductImage(libs.DB, payload.Name, orientation)
	return nil
}

func decodeBase64(encodedStr string) ([]byte, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedStr)
	if err != nil {
		return nil, err
	}
	return decodedBytes, nil
}
