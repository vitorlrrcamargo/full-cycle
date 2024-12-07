package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	fmt.Println(len(myArray))            // array size
	fmt.Println(len(myArray) - 1)        // last position of the array
	fmt.Println(myArray[0])              // value of the first position of the array
	fmt.Println(myArray[len(myArray)-1]) // value of the last position of the array

	// traversing the array
	for i, v := range myArray {
		println()
		fmt.Printf("The index value is %d and the value is %d", i, v) // %d - integer (base 10)
	}
}
