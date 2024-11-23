package main

type Conta struct {
	saldo int
}

func newConta(valor int) *Conta {
	return &Conta{saldo: valor}
}

func (c Conta) simularDeposito(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func (c *Conta) depositar(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

func main() {
	conta := newConta(100)
	conta.simularDeposito(200)
	println(conta.saldo)
	conta.depositar(400)
	println(conta.saldo)
}
