package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/libs"
	"github.com/fusionxx23/ecommerce-go/models"
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
