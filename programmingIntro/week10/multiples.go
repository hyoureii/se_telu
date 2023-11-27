package main

import (
	"fmt"
)

func main() {
	var num uint
	fmt.Printf("Enter a positive integer\n")
	fmt.Scan(&num)
	if num%5 == 0 {
		fmt.Printf("Kelipatan 5\n")
	}
	if num%3 == 0 {
		fmt.Printf("Kelipatan 3\n")
	}
}
