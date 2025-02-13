package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("Received request for: ", r.URL.Path)
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	log.Print("Starting server on port 8080")
	err := http.ListenAndServe(":8080", nil) // Changed port from 80 to 8080
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
