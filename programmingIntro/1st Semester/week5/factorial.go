package main

import (
	"fmt"
)

func factorial(num uint) uint {
	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}

func main() {
	var x uint
	fmt.Printf("Enter a positive integer\n")
	fmt.Scan(&x)
	fmt.Printf("%d! = %d", x, factorial(x))
}
