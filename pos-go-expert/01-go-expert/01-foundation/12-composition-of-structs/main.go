package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Customer struct {
	Name   string
	Age    int
	Active bool
	Address
}

func main() {
	vitor := Customer{
		Name:   "Vitor",
		Age:    33,
		Active: true,
	}
	vitor.Active = false
	vitor.City = "São Paulo"
	fmt.Printf("Name: %s, Age: %d, Active: %t, City: %s", vitor.Name, vitor.Age, vitor.Active, vitor.City) // %t - bool
}
