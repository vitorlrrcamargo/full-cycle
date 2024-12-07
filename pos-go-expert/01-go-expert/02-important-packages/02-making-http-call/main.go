package main

import (
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.google.com")
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(data))
	resp.Body.Close()
}
