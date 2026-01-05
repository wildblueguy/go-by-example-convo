package main

import "fmt"

func sumAndCount(vals ...int) (int, int) {
  sum := 0
  for i := 0; i < len(vals); i++ {
    sum += vals[i]
  }
  return sum, len(vals)
}

func main() {
  fmt.Println(sumAndCount(1, 2, 3, 4))
}
