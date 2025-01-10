package main

import "github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/09-handling-events/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	rabbitmq.Publish(ch, "Hello, RabbitMQ!", "amq.direct")
}
