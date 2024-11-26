package main

import "fmt"

func main() {
	fmt.Println("Primeira Linha")
	defer fmt.Println("Segunda Linha") // última execução
	fmt.Println("Terceira Linha")
	defer fmt.Println("Quarta Linha") // penúltima execução
	fmt.Println("Quinta Linha")
}
