package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"` // `json:"-"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	jsonData, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}
	println(string(jsonData))

	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"n":2,"s":200}`)
	var conta2 Conta
	err = json.Unmarshal(jsonPuro, &conta2)
	if err != nil {
		panic(err)
	}
	println(conta2.Numero)
	println(conta2.Saldo)
}
