package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int64
	var syncPoint sync.WaitGroup
	syncPoint.Add(50) // Not necessary with syncPoint.Go
	for i := 0; i < 50; i++ {
		go func() { // Alternative: `syncPoint.Go(func() {...})`
			counter.Add(1000)
			syncPoint.Done() // Not necessary with syncPoint.Go
		}()
		
	}
	syncPoint.Wait()
	fmt.Println(counter.Load())
}
