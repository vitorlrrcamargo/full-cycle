package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("The value of (a) is %v and the memory address is %v\n", a, &a)
	var b *int = &a
	fmt.Printf("The value of (b) is %v, the memory address is %v and the pointer is for the value %v\n\n", b, &b, *b)
	*b = 20
	fmt.Printf("The value of (a) is %v and the memory address is %v\n", a, &a)
	fmt.Printf("The value of (b) is %v, the memory address is %v and the pointer is for the value %v\n", b, &b, *b)
}
