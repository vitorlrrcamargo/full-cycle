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
	case <-time.After(time.Second * 5):
		// print on command line (stdout)
		log.Println("Request processed successfully")
		// print on browser
		w.Write([]byte("Request processed successfully"))
	case <-ctx.Done():
		// print on command line (stdout)
		log.Println("Request canceled by client")
	}
}
