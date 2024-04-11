package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Rabbit MQ")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("TestQ", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(q)
	err = ch.Publish("", "TestQueue", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello World"),
	},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully Published Message to Queue")

}
