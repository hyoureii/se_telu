package main

import (
  "fmt"
)

func main() {
  var squares int
  fmt.Println("Input how much squares and the side length of each square")
  fmt.Scan(&squares)
  sides := make([]int, squares)
  for i := 0; i < squares; i++ {
    fmt.Scan(&sides[i])
    fmt.Printf("%d %d\n", sides[i]*sides[i], sides[i]*4)
  }
}