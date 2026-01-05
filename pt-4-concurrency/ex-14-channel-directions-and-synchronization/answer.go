package main

import "fmt"

func main() {
	channel := make(chan string)
	go func() {
		var sendSide chan<- string = channel
		sendSide <- "Hello"
	}()
	var receiveSide <-chan string = channel
	message := <-receiveSide
	fmt.Println(message)
}
