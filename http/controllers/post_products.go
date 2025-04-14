package controllers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"io"
	"net/http"
	"strconv"

	"github.com/chai2010/webp"
	"github.com/fusionxx23/ecommerce-go/http/libs"
	"github.com/fusionxx23/ecommerce-go/http/models"
	"github.com/streadway/amqp"
)

func postProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err = models.InsertProduct(&product)
	if err != nil {
		http.Error(w, "Error inserting product.", http.StatusBadRequest)
		return
	}
	//status 200
	w.WriteHeader(http.StatusOK)
}

func postProductVariant(w http.ResponseWriter, r *http.Request) {
	var productVariant models.ProductVariant
	err := json.NewDecoder(r.Body).Decode(&productVariant)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Println(productVariant.ProductID)
	err = models.InsertProductVariant(&productVariant)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error inserting product variant.", http.StatusBadRequest)
		return
	}
}

func postProductImage(w http.ResponseWriter, r *http.Request) {

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // Limit upload size to 10MB
	if err != nil {
		http.Error(w, "Form data to big", http.StatusBadRequest)
		return
	}

	imageFile, _, err := r.FormFile("image_file")
	if err != nil {
		http.Error(w, "Unable to get image file", http.StatusBadRequest)
		return
	}
	defer imageFile.Close()

	// Create a bytes.Buffer
	var buffer bytes.Buffer

	// Copy the file content into the buffer
	_, err = io.Copy(&buffer, imageFile)
	if err != nil {
		http.Error(w, "Unable to copy image file", http.StatusInternalServerError)
		return
	}
	// Check if the file is a valid image
	img, _, err := image.Decode(&buffer)
	if err != nil {
		http.Error(w, "Uploaded file is not a valid image", http.StatusBadRequest)
		return
	}

	// Convert the image to WebP format
	var webpBuffer bytes.Buffer
	err = webp.Encode(&webpBuffer, img, &webp.Options{Lossless: true})
	if err != nil {
		http.Error(w, "Failed to convert image to WebP", http.StatusInternalServerError)
		return
	}

	product_id := r.FormValue("product_id") // Replace "fieldName" with your form field name
	productID, err := strconv.ParseInt(product_id, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	fmt.Println(productID, "Product ID")
	p := models.ProductImage{ProductID: productID}
	models.InsertProductImage(&p)
	err = libs.RabbitChannel.Publish("", "ImageQueue", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body: fmt.Appendf(nil, `{"name": "%d", "bytes": "%s"}`, p.ImageId,
			base64.StdEncoding.EncodeToString(webpBuffer.Bytes())),
	})
	if err != nil {
		fmt.Println("Failed to publish a message:", err)
	}

	fmt.Println("Sent Message")
	w.Write([]byte("Form data processed successfully"))
}
