package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Vitor"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("The type of E is %T", e) // variable type
	println()
	fmt.Printf("The value of E is %v", e) // variable type
}
