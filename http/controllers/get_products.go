package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fusionxx23/ecommerce-go/http/libs"
	"github.com/fusionxx23/ecommerce-go/http/models"
	"github.com/gorilla/mux"
)

func getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := models.SelectAllProducts()
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	productsJSON, err := json.Marshal(products)
	if err != nil {
		http.Error(w, "Failed to marshal json", http.StatusInternalServerError)
		return
	}
	libs.SendJson(w, productsJSON)
}
func getProductFromSlug(w http.ResponseWriter, r *http.Request) {
	// get search query from url
	vars := mux.Vars(r)
	s, exists := vars["slug"]
	if !exists {
		http.Error(w, "Error getting slug", http.StatusInternalServerError)
		return
	}
	product, err := models.SelectProductFromSlug(s)
	if err != nil {
		http.Error(w, "Error  product", http.StatusInternalServerError)
		return
	}
	productJSON, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Failed to marshal json", http.StatusInternalServerError)
		return
	}
	libs.SendJson(w, productJSON)
}

func getProductVariants(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	i, exists := vars["id"]
	if !exists {
		http.Error(w, "Error getting id", http.StatusInternalServerError)
		return
	}
	// convert string to int
	id, err := strconv.Atoi(i)
	if err != nil {
		http.Error(w, "Error converting id", http.StatusInternalServerError)
		return
	}
	// get search query from url
	variants, err := models.SelectProductVariants(int64(id))
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error getting product variants", http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(variants)
	if err != nil {
		http.Error(w, "Failed to marshal json", http.StatusInternalServerError)
	}
	libs.SendJson(w, j)
}
