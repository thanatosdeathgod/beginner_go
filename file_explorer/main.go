package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Get the directory from the command line
	directory := os.Args[1]

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <directory>")
		os.Exit(1)
	}

	// Walk the directory
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
