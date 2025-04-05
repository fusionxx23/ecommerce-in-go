package controllers

import (
	"fmt"
	"github.com/fusionxx23/ecommerce-go/libs"
	"github.com/gorilla/mux"
	"net/http"
)

func CheckoutHandler(s *mux.Router) {
	s.HandleFunc("", getCheckout).Methods("GET")
}

func getCheckout(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get the cart
	cartCookie, err := r.Cookie("cart")
	if err != nil {
		id, err := libs.GenerateRandomID(32)
		if err != nil {
			http.Error(w, "Error generating cart ID", http.StatusInternalServerError)
			return
		}
		// Create a new cart cookie if it doesn't exist
		cartCookie = &http.Cookie{
			Name:  "cart",
			Value: id, // Replace with actual cart initialization logic
			Path:  "/",
		}
	}
	fmt.Fprintf(w, "Cart: %s", cartCookie.Value)
}
