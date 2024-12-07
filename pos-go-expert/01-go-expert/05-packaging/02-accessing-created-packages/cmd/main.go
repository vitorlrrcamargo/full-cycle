package main

import (
	"fmt"

	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/05-packaging/02-accessing-created-packages/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
