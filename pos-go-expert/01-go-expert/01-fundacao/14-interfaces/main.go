package main

import "fmt"

type Pessoa interface {
	GetNome() string
	Desativar(nome string)
}

type Funcionario struct {
	Nome string
}

type Cliente struct {
	Nome string
}

func (f Funcionario) GetNome() string {
	return f.Nome
}

func (f Funcionario) Desativar(nome string) {
	fmt.Printf("O funcionario %s foi desativado", nome)
}

func (c Cliente) GetNome() string {
	return c.Nome
}

func (c Cliente) Desativar(nome string) {
	fmt.Printf("O cliente %s foi desativado", nome)
}

func Desativacao(pessoa Pessoa) {
	fmt.Printf("Desativando pessoa... ")
	pessoa.Desativar(pessoa.GetNome())
}

func main() {
	vitor := Funcionario{
		Nome: "Vitor Camargo",
	}

	joao := Cliente{
		Nome: "João Silva",
	}

	Desativacao(vitor)
	println()
	Desativacao(joao)
}
