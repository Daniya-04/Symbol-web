package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Define what happens when someone visits your website's home page ("/")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! I built this web server from scratch using Go!")
	})

	// 2. Start the server on port 8080 and tell it to listen
	fmt.Println("Server is running! Open http://localhost:8080 in your browser.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
