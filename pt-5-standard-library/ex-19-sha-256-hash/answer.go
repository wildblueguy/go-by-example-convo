package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	hash := sha256.Sum256([]byte("gobyexample"))
	fmt.Println(hex.EncodeToString(hash[:]))

	// Streaming alternative:
	// hasher := sha256.New()
	// hasher.Write([]byte("gobyexample"))
	// hash := hasher.Sum(nil)
	// fmt.Println(hex.EncodeToString(hash))
}
