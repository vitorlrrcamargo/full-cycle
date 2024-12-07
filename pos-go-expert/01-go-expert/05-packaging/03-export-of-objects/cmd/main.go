package main

import (
	"fmt"

	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/05-packaging/03-export-of-objects/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 5
	fmt.Println(m.Add())
	fmt.Println(math.X)
	fmt.Println(m.C)
}
