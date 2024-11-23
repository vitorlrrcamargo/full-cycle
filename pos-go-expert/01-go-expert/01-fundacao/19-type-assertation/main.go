package main

import "fmt"

func main() {
	var minhaVar interface{} = "Vitor Camargo"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é '%v' e o resultado de ok é '%v'", res, ok)
}
