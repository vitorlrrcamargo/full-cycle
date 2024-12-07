package main

import (
	"fmt"

	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/01-foundation/21-packages-and-modules/mathematics"
)

func main() {
	s := mathematics.Sum(10, 20)
	c := mathematics.Car{Brand: "Fiat"}

	fmt.Println(c.Brand)
	fmt.Println(c.Run())
	fmt.Println("Result:", s)
	fmt.Println(mathematics.A)
}
