package initializers

import (
	"fmt"

	"github.com/fusionxx23/ecommerce-go/http/libs"
	"github.com/streadway/amqp"
)

func InitRabbitMQ() {
	fmt.Println("Go RabbitMQ Tutorial")

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
	libs.RabbitChannel = c
	_, err = c.QueueDeclare("ImageQueue", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
