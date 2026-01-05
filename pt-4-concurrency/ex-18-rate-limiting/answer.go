package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	requestChannel := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requestChannel <- i
	}
	close(requestChannel)
	tickChannel := time.Tick(200 * time.Millisecond)
	for {
		select {
		case <-tickChannel:
			fmt.Println(<-requestChannel)
			if len(requestChannel) == 0 {
				os.Exit(0)
			}
		}
	}
}
