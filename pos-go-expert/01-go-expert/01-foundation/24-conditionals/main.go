package main

func main() {
	a := 3
	b := 2
	c := 1

	if a > b {
		println("a > b")
	} else if a < b {
		println("a < b")
	} else {
		println("a == b")
	}

	if a > b && a > c {
		println("a > b && a > c")
	}

	if a > b || a > c {
		println("a > b || a > c")
	}

	switch a {
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	default:
		println("other")
	}
}
