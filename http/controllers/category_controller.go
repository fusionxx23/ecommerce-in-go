package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/http/libs"
	"github.com/fusionxx23/ecommerce-go/http/models"
	"github.com/gorilla/mux"
)

func CategoriesHandler(s *mux.Router) {
	s.HandleFunc("", getCategories).Methods("GET")
	s.HandleFunc("", createCategory).Methods("POST")
	s.HandleFunc("/delete", deleteCategory).Methods("PUT")
	s.HandleFunc("/add-product-category", addProductToCategory).Methods("POST")
}

func getCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := models.SelectCategories()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	p, err := json.Marshal(categories)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	libs.SendJson(w, p)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	var payload models.Category
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	err = models.InsertCategory(&payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		CategoryID int64 `json:"categoryId"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = models.DeleteCategory(payload.CategoryID)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type ProductRef struct {
	ProductID  int64 `json:"productId"`
	CategoryID int64 `json:"categoryId"`
}

func addProductToCategory(w http.ResponseWriter, r *http.Request) {
	var p ProductRef
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	resp := models.ProductCategoryRef{ProductId: p.ProductID, CategoryId: p.CategoryID}
	err = models.InsertProductCategoryRef(&resp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
