package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("O valor de (a) é %v e o endereço da memória é %v\n", a, &a)
	var b *int = &a
	fmt.Printf("O valor de (b) é %v, o endereço da memória é %v e o apontamento é para o valor %v\n\n", b, &b, *b)
	*b = 20
	fmt.Printf("O valor de (a) é %v e o endereço da memória é %v\n", a, &a)
	fmt.Printf("O valor de (b) é %v, o endereço da memória é %v e o apontamento é para o valor %v\n", b, &b, *b)
}
