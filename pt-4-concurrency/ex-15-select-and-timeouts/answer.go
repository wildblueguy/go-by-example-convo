package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Hello"
	}()
	select {
	case message := <-channel:
		fmt.Println(message)
	case <-time.After(time.Second):
		fmt.Println("done")
	}
}
