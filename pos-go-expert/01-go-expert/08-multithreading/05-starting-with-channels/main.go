package main

import "fmt"

// thread 1
func main() {
	channel := make(chan string) // empty

	// thread 2
	go func() {
		channel <- "Hello World!" // full
	}()

	// thread 1
	msg := <-channel // empty
	fmt.Println(msg)
}
