package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	vitor := Cliente{
		Nome:  "Vitor",
		Idade: 33,
		Ativo: true,
	}
	vitor.Ativo = false

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t", vitor.Nome, vitor.Idade, vitor.Ativo) // %t - bool
}
