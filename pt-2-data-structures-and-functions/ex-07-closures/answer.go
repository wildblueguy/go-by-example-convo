package main

import "fmt"

func newSequence() (next func() int) {
  i := 0
  next = func() int {
      i++
      return i
  }
  return
}

func main() {
  next := newSequence()
  fmt.Println(next())
  fmt.Println(next())
  fmt.Println(next())
}
