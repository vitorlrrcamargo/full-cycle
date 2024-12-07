package main

import (
	"github.com/google/uuid"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/05-packaging/06-using-workspaces/math"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
	println(uuid.New().String())
}
