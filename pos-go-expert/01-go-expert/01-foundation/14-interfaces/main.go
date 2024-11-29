package main

import "fmt"

type Person interface {
	GetName() string
	Disable(Name string)
}

type Employee struct {
	Name string
}

type Customer struct {
	Name string
}

func (f Employee) GetName() string {
	return f.Name
}

func (f Employee) Disable(Name string) {
	fmt.Printf("The Employee %s has been deactivated", Name)
}

func (c Customer) GetName() string {
	return c.Name
}

func (c Customer) Disable(Name string) {
	fmt.Printf("The Customer %s has been deactivated", Name)
}

func Deactivation(Person Person) {
	fmt.Printf("Disabling Person... ")
	Person.Disable(Person.GetName())
}

func main() {
	vitor := Employee{
		Name: "Vitor Camargo",
	}

	joao := Customer{
		Name: "João Silva",
	}

	Deactivation(vitor)
	println()
	Deactivation(joao)
}
