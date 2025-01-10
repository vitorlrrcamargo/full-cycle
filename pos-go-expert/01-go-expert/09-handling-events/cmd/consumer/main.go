package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/09-handling-events/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs, "my-queue")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
