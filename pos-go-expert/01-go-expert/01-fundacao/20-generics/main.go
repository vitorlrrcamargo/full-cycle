package main

import "fmt"

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Vitor": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Vitor": 1000.10, "João": 2000.15, "Maria": 3000.20}
	m3 := map[string]MyNumber{"Vitor": 1000, "João": 2000, "Maria": 3000}

	fmt.Printf("%v\n", Soma(m))
	fmt.Printf("%v\n", Soma(m2))
	fmt.Printf("%v\n", Soma(m3))

	fmt.Printf("%v\n", Compara(10, 10.0))
}
