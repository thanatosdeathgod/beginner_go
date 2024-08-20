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

	// Handle all requests with the file server
	http.Handle("/", fs)

	// Start the server
	port := os.Args[2]
	fmt.Printf("Starting server on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
