package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id          int64
	Description string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var id int64 = 1

	// RabbitMQ
	go func() {
		for {
			atomic.AddInt64(&id, 1)
			time.Sleep(time.Second * 1)
			message := Message{id, "Hello from RabbitMQ"}
			c1 <- message
		}
	}()

	// Kafka
	go func() {
		for {
			atomic.AddInt64(&id, 1)
			time.Sleep(time.Second * 2)
			message := Message{id, "Hello from Kafka"}
			c2 <- message
		}
	}()

	//for i := 0; i < 2; i++ {
	for {
		select {
		case message := <-c1:
			fmt.Printf("Received from RabbitMQ: ID: %d - %s\n", message.Id, message.Description)

		case message := <-c2:
			fmt.Printf("Received from Kafka: ID: %d - %s\n", message.Id, message.Description)

		case <-time.After(time.Second * 3):
			println("timeout")

			//default:
			//	println("default")
		}
	}
}
