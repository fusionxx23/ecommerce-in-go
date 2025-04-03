package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

func AuthHandler(s *mux.Router) {
	s.HandleFunc("/{provider}", signInWithGoogle)
	s.HandleFunc("/{provider}/callback", callback)
}

type key int

// ProviderParamKey can be used as a key in context when passing in a provider
const ProviderParamKey key = iota

func signInWithGoogle(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func callback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error completing user auth", http.StatusInternalServerError)
		return
	}
	fmt.Println(user)
	http.Redirect(w, r, "/", http.StatusFound)
}
