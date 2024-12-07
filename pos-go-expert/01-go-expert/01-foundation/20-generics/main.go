package main

import "fmt"

type MyNumber int

type Number interface {
	~int | float64
}

func Sum[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Vitor": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Vitor": 1000.10, "João": 2000.15, "Maria": 3000.20}
	m3 := map[string]MyNumber{"Vitor": 1000, "João": 2000, "Maria": 3000}

	fmt.Printf("%v\n", Sum(m))
	fmt.Printf("%v\n", Sum(m2))
	fmt.Printf("%v\n", Sum(m3))

	fmt.Printf("%v\n", Compare(10, 10.0))
}
