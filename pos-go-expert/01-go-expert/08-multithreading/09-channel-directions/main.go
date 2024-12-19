package main

import "fmt"

// thread 1
func main() {
	hello := make(chan string)

	// thread 2
	go receive("Hello", hello)
	send(hello)
}

func receive(name string, hello chan<- string) {
	hello <- name
}

func send(data <-chan string) {
	fmt.Println(<-data)
}
