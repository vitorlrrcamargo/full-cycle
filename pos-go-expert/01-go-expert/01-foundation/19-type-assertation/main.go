package main

import "fmt"

func main() {
	var myVar interface{} = "Vitor Camargo"
	println(myVar.(string))
	res, ok := myVar.(int)
	fmt.Printf("The (res) value is '%v' and the result of (ok) is '%v'", res, ok)
}
