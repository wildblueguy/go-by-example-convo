package main

import (
	"fmt"
	"math"
)

type Circle struct {
	r float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	circle := Circle{r: 10}
	fmt.Println(circle.Area())
}
