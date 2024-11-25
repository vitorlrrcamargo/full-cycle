package main

import (
	"curso-go/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	c := matematica.Carro{Marca: "Fiat"}

	fmt.Println(c.Marca)
	fmt.Println(c.Andar())
	fmt.Println("Resultado:", s)
	fmt.Println(matematica.A)
}
