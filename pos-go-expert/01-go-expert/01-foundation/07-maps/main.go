package main

import "fmt"

func main() {
	// salaries := make(map[string]int)
	// salaries := map[string]int{}
	salaries := map[string]int{"Vitor": 1000, "João": 2000, "Maria": 3000}

	fmt.Println(salaries["Vitor"])
	delete(salaries, "Vitor")
	salaries["Pedro"] = 4000
	fmt.Println(salaries["Pedro"])

	for nome, salario := range salaries {
		fmt.Printf("\nThe salary of %s is %d", nome, salario) // %s - string
	}
}
