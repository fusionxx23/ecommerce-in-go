package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/database"
	"github.com/fusionxx23/ecommerce-go/models"
	"github.com/markbates/goth/gothic"
)

func GetCart(w http.ResponseWriter, r *http.Request) {
	// get gothic session
	a, err := gothic.GetFromSession("email", r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	// Implement logic to get the cart
	cartCookie, err := r.Cookie("cart")
	if err != nil {

		c, err := createCart(w)
		if err != nil {
			return
		}
		SendCartJson(c, w)
		return
	}

	var cart models.Cart
	// get cart with cart id
	result := database.DB.First(&cart, "id = ?", cartCookie.Value)
	if result.Error != nil {
		cart, err := createCart(w)
		if err != nil {
			return
		}
		SendCartJson(cart, w)
		return
		// cart not found
	} else {
		SendCartJson(cart, w)
		return
	}
}

func SendCartJson(cart models.Cart, w http.ResponseWriter) {
	cartJSON, err := json.Marshal(cart)
	if err != nil {
		http.Error(w, "Error marshaling cart to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cartJSON)
}

func createCart(w http.ResponseWriter) (models.Cart, error) {
	c := models.Cart{}
	// insert a cart row
	result := database.DB.Create(&c)
	if result.Error != nil {
		fmt.Println(result.Error)
		// respond with 405
		http.Error(w, "Error creating cart", http.StatusInternalServerError)
	}
	return c, result.Error
}
