package controllers

import (
	"github.com/gorilla/mux"
)

func CartHandler(s *mux.Router) {
	s.HandleFunc("", GetCart).Methods("GET")
	s.HandleFunc("/items", CartItem).Methods("POST", "DELETE")
}
