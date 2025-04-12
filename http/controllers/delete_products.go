package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/http/models"
)

type RequestBody struct {
	ID int64 `json:"id"`
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	var body RequestBody

	// Parse the JSON body
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusBadRequest)
		return
	}
	err = models.DeleteProduct(body.ID)
	if err != nil {
		http.Error(w, "Error deleting product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
