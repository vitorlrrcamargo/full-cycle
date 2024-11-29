package main

import "net/http"

func main() {
	http.HandleFunc("/", ZipCodeSearch)
	http.ListenAndServe(":8080", nil)
}

func ZipCodeSearch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
