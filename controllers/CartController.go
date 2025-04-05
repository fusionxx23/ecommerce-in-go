package controllers

import (
	cart "github.com/fusionxx23/ecommerce-go/controllers/cart"
	"github.com/gorilla/mux"
)

func CartHandler(s *mux.Router) {
	s.HandleFunc("", cart.GetCart).Methods("GET")
	s.HandleFunc("/items", cart.CartItem).Methods("POST", "DELETE")
}
