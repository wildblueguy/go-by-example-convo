package main

import "fmt"

func main() {
	var integer int
	pointer := &integer
	*pointer = 1
	fmt.Println(integer, *pointer)
}
