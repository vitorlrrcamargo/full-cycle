package main

func soma(a, b *int) int {
	*a = 30
	*b = 40
	return *a + *b
}

func main() {
	myVar1 := 10
	myVar2 := 20
	soma(&myVar1, &myVar2)
	println(myVar1)
	println(myVar2)
}
