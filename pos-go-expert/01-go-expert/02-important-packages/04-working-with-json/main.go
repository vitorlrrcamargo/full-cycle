package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int `json:"n"` // `json:"-"`
	Balance int `json:"s"`
}

func main() {
	account := Account{Number: 1, Balance: 100}
	jsonData, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	println(string(jsonData))

	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	jsonData = []byte(`{"n":2,"s":200}`)
	var account2 Account
	err = json.Unmarshal(jsonData, &account2)
	if err != nil {
		panic(err)
	}
	println(account2.Number)
	println(account2.Balance)
}
