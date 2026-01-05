package main

import "fmt"

type Describer interface {
	Description() string
}

type InnerDescriber struct{}

func (id InnerDescriber) Description() string {
	return "Inner description"
}

type OuterDescriber struct {
	InnerDescriber
}

func main() {
	var describer Describer = OuterDescriber{}
	fmt.Println(describer.Description())
}
