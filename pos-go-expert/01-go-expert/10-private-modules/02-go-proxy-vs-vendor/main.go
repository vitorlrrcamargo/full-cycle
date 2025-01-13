package main

import (
	"fmt"

	"github.com/vitorlrrcamargo/full-cycle-private/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
