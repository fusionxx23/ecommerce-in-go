package main

import (
	"log"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/http/controllers"
	"github.com/fusionxx23/ecommerce-go/http/initializers"
	"github.com/gorilla/mux"
)

func init() {
	initializers.ConnectDatabase()
	initializers.SyncDb()
	initializers.InitOAuth()
	initializers.InitRabbitMQ()
}

func main() {
	r := mux.NewRouter()
	auth := r.PathPrefix("/auth").Subrouter()
	cart := r.PathPrefix("/cart").Subrouter()
	products := r.PathPrefix("/products").Subrouter()
	controllers.AuthHandler(auth)
	controllers.CartHandler(cart)
	controllers.ProductController(products)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
