package main

import (
  "fmt"
)

func main() {
  var amount uint
  var text string
  fmt.Println("Input the amount to print and the String")
  fmt.Scan(&amount, &text)
  for i := 0; i < amount; i++ {
    fmt.Println(text)
  }
}