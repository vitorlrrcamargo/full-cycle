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

func (c *Customer) Desativar() {
	if c.Active == true {
		c.Active = false
		fmt.Printf("The customer %s has been deactivated\n", c.Name)
	}
}

func main() {
	vitor := Customer{
		Name:   "Vitor",
		Age:    33,
		Active: true,
	}
	vitor.Desativar()
	vitor.City = "São Paulo"
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s", vitor.Name, vitor.Age, vitor.Active, vitor.City) // %t - bool
}
