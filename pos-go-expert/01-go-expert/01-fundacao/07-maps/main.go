package main

import "fmt"

func main() {
	// salarios := make(map[string]int)
	// salarios := map[string]int{}
	salarios := map[string]int{"Vitor": 1000, "João": 2000, "Maria": 3000}

	fmt.Println(salarios["Vitor"])
	delete(salarios, "Vitor")
	salarios["Pedro"] = 4000
	fmt.Println(salarios["Pedro"])

	for nome, salario := range salarios {
		fmt.Printf("\nO salário de %s é %d", nome, salario) // %s - string
	}
}
