package controllers

import (
	"github.com/gorilla/mux"
)

func ProductController(s *mux.Router) {
	s.HandleFunc("", getProducts).Methods("GET")
	s.HandleFunc("/{slug}", getProductFromSlug).Methods("GET")
	s.HandleFunc("", postProduct).Methods("POST")
}
