package main

import (
	"course-go/mathematics"
	"fmt"
)

func main() {
	s := mathematics.Sum(10, 20)
	c := mathematics.Car{Brand: "Fiat"}

	fmt.Println(c.Brand)
	fmt.Println(c.Run())
	fmt.Println("Result:", s)
	fmt.Println(mathematics.A)
}
