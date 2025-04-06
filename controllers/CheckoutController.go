package controllers

import (
	"github.com/gorilla/mux"
	"net/http"
)

func CheckoutHandler(s *mux.Router) {
	s.HandleFunc("", getCheckout).Methods("GET")
}

func getCheckout(w http.ResponseWriter, r *http.Request) {
	// Implement logic to get the cart
}
