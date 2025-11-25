package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, e *http.Request) {
		w.Write([]byte("OK"))
	})
		
	mux.HandleFunc("/hello", func(w http.ResponseWriter, e *http.Request) {
		w.Write([]byte("Hello from Themis!"))
	})

	log.Println("Themis listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}