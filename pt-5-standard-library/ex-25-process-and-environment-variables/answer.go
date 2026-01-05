package main

import (
	"fmt"
	"os"
)

func main() {
	value, isSet := os.LookupEnv("HOSTNAME") // Linux convention
	if isSet {
		fmt.Println(value)
	} else {
		fmt.Println("not set")
	}
}
