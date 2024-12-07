package main

type Account struct {
	balance int
}

func newAccount(value int) *Account {
	return &Account{balance: value}
}

func (c Account) simulateDeposit(value int) int {
	c.balance += value
	println(c.balance)
	return c.balance
}

func (c *Account) deposit(value int) int {
	c.balance += value
	println(c.balance)
	return c.balance
}

func main() {
	Account := newAccount(100)
	Account.simulateDeposit(200)
	println(Account.balance)
	Account.deposit(400)
	println(Account.balance)
}
