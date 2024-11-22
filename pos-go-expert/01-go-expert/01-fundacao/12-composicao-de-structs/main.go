package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	vitor := Cliente{
		Nome:  "Vitor",
		Idade: 33,
		Ativo: true,
	}
	vitor.Ativo = false
	vitor.Cidade = "São Paulo"
	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t, Cidade: %s", vitor.Nome, vitor.Idade, vitor.Ativo, vitor.Cidade) // %t - bool
}
