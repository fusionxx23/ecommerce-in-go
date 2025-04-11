package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/models"
)

func postProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
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
func postProductVariant() {

}
