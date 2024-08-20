package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <file> <key>")
		os.Exit(1)
	}

	file := os.Args[1]
	key := os.Args[2]

	// Read the file
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Create the cipher
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println("Error creating cipher:", err)
		os.Exit(1)
	}

	// Create the GCM
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		fmt.Println("Error creating GCM:", err)
		os.Exit(1)
	}

	// Create a nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		fmt.Println("Error creating nonce:", err)
		os.Exit(1)
	}

	// Encrypt the data of file specified
	encrypted := gcm.Seal(nonce, nonce, data, nil)
	os.WriteFile(file+".enc", encrypted, 0644)
}
