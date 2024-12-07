package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file creation

	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	// size, err := f.WriteString("Hello, World!")
	size, err := file.Write([]byte("Writing data to the file!"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created successfully! Size: %d bytes\n", size)

	file.Close()

	// file reading

	file2, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file2))

	// reading little by little opening the file

	file3, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file3)
	buffer := make([]byte, 5)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	file3.Close()

	// remoção  de file

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
