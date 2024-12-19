package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}

// thread 1
func main() {
	data := make(chan int)
	workersQuantity := 100

	for i := 0; i < workersQuantity; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1000; i++ {
		data <- i
	}
}
