package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Get the directory from the command line
	directory := os.Args[1]

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <directory> <port>")
		os.Exit(1)
	}

	// Create a file server
	fs := http.FileServer(http.Dir(directory))
	// FileServer = FileServer is the function to create the file server
	// Dir = Dir means the function to get the directory

	// Handle all requests with the file server
	http.Handle("/", fs)
	// Handle is the function to handle the request

	// Start the server
	port := os.Args[2]
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	// ListenAndServe means the function to start the server
	// ":" means the port number, nil means the handler
}
