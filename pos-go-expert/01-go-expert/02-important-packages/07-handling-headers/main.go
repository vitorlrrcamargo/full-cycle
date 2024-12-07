package main

import "net/http"

func main() {
	http.HandleFunc("/", ZipCodeSearchHandler)
	http.ListenAndServe(":8080", nil)
}

func ZipCodeSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	zipCodeParam := r.URL.Query().Get("zipcode")
	if zipCodeParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
