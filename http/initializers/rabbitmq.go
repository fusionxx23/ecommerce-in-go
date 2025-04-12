package initializers

import (
	"fmt"

	"github.com/fusionxx23/ecommerce-go/http/libs"
	"github.com/streadway/amqp"
)

func InitRabbitMQ() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	libs.RabbitMqConn = conn
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
}
