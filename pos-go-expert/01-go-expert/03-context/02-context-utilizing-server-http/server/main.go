package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request started")
	defer log.Println("Request completed")
	select {
	case <-time.After(5 * time.Second):
		// Print to command line stdout
		log.Println("Request processed successfully")
		// Print in browser
		w.Write([]byte("Request processed successfully"))
	case <-ctx.Done():
		// Print to command line stdout
		log.Println("Request canceled by client")
	}
}
