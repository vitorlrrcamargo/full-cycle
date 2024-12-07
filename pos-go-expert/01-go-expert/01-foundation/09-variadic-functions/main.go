package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(20, 30, 40, 50, 60))
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
