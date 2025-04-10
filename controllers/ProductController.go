package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/libs"
	"github.com/fusionxx23/ecommerce-go/models"
	"github.com/gorilla/mux"
)

func ProductController(s *mux.Router) {
	s.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
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
	}).Methods("GET")
	s.HandleFunc("/{slug}", GetProductFromSlug).Methods("GET")
}

func GetProductFromSlug(w http.ResponseWriter, r *http.Request) {
	// get search query from url
	vars := mux.Vars(r)
	s, exists := vars["slug"]
	if !exists {
		http.Error(w, "Error getting slug", http.StatusInternalServerError)
		return
	}
	product, err := models.SelectProductFromSlug(s)
	if err != nil {
		http.Error(w, "Error retrieving product", http.StatusInternalServerError)
		return
	}
	productJSON, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Failed to marshal json", http.StatusInternalServerError)
		return
	}
	libs.SendJson(w, productJSON)
}
