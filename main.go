package main

import (
	"log"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/controllers"
	"github.com/fusionxx23/ecommerce-go/initializers"
	"github.com/gorilla/mux"
)

func init() {
	initializers.ConnectDatabase()
	initializers.SyncDb()
	initializers.InitOAuth()
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/auth").Subrouter()
	controllers.AuthHandler(s)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
