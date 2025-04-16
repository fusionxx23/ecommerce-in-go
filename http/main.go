package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

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

	// get RABBITMQ_HOST from env
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	// check if rabbitmqHost is empty
	if rabbitmqHost == "" {
		rabbitmqHost = "localhost"
	}
	connUrl := fmt.Sprintf("amqp://guest:guest@%s:5672/", rabbitmqHost)
	// create RabbitMQ connection and create channel
	conn, err := amqp.Dial(connUrl)
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
	categories := r.PathPrefix("/categories").Subrouter()
	products := r.PathPrefix("/products").Subrouter()
	controllers.AuthHandler(auth)
	controllers.CartHandler(cart)
	controllers.ProductController(products)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))

}
