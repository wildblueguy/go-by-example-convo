package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(max(1, 2))
}
