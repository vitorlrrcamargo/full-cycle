package main

import "github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/05-packaging/05-working-with-go-mod-replace/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}
