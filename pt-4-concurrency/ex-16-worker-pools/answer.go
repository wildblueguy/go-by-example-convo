package main

import "fmt"

func doWork(jobChannel <-chan int, resultChannel chan<- int) {
	for job := range jobChannel { // Closed elsewhere
		resultChannel <- job * 2
	}
}

func main() {
	// At least one of the channels must be buffered
	jobChannel := make(chan int, 30)
	resultChannel := make(chan int)
	for i := 0; i < 3; i++ {
		go doWork(jobChannel, resultChannel)
	}
	for i := 0; i < 30; i++ {
		jobChannel <- i
	}
	close(jobChannel)         // Ranged-over elsewhere
	for i := 0; i < 30; i++ { // Can't use `range resultChannel` (no dedicated closer)
		fmt.Println(<-resultChannel)
	}
}
