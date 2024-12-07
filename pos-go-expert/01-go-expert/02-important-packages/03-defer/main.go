package main

import "fmt"

func main() {
	fmt.Println("First Line")
	defer fmt.Println("Second Line") // last execution
	fmt.Println("Third Line")
	defer fmt.Println("Fourth Line") // penultimate execution
	fmt.Println("Fifth Line")
}
