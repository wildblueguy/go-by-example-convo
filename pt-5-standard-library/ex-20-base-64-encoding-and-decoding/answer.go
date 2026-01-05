package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	encoded := base64.StdEncoding.EncodeToString([]byte("golang"))
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(encoded, string(decoded))
}
