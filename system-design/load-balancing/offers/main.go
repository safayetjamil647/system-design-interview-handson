package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define a handler function to handle incoming HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Write the response to the client
		fmt.Fprintf(w, "Hello, World!")
	})

	// Start the web server on port 8080
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
