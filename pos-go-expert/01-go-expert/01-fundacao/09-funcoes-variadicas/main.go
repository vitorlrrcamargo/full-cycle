package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(20, 30, 40, 50, 60))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
