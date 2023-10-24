package main

import (
	"fmt"
)

func main() {
	var num, total uint
  fmt.Println("Input however many numbers you like")
	for true {
		fmt.Scan(&num)
		if num%2 == 0 {
			total += num
		} else {
      break
    }
	}
  fmt.Printf("Total : %d", total)
}
