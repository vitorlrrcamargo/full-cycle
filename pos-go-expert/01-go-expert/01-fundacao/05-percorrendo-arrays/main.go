package main

import "fmt"

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(len(meuArray))             // tamanho do array
	fmt.Println(len(meuArray) - 1)         // última posição do array
	fmt.Println(meuArray[0])               // valor da primeira posição do array
	fmt.Println(meuArray[len(meuArray)-1]) // valor da última posição do array

	// percorrendo o array
	for i, v := range meuArray {
		println()
		fmt.Printf("O valor do índice é %d e o valor é %d", i, v) // %d - integer (base 10)
	}
}
