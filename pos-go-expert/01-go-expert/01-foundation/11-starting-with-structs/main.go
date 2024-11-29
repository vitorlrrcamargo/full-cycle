package main

import "fmt"

type Customer struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	vitor := Customer{
		Name:   "Vitor",
		Age:    33,
		Active: true,
	}
	vitor.Active = false

	fmt.Printf("Name: %s, Age: %d, Active: %t", vitor.Name, vitor.Age, vitor.Active) // %t - bool
}
