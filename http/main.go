package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fusionxx23/ecommerce-go/http/controllers"
	"github.com/fusionxx23/ecommerce-go/http/initializers"
	"github.com/fusionxx23/ecommerce-go/http/libs"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
)

func init() {
	initializers.ConnectDatabase()
	initializers.SyncDb()
	initializers.InitOAuth()
}

func main() {
	// create RabbitMQ connection and create channel
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}
	libs.RabbitMqConn = conn
	defer conn.Close()

	fmt.Println("Succesfully connected to RabbitMQ")
	c, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer c.Close()
	libs.RabbitChannel = c
	_, err = c.QueueDeclare("ImageQueue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

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
