package main

import (
	"fmt"
)

func chopDigits(num uint) {
	var current, total uint
	for num > 0 {
		current = num % 10
		num /= 10
		total += current
	}
	fmt.Printf("Total : %d", total)
}

func main() {
	var x uint
	fmt.Printf("Enter a positive integer\n")
	fmt.Scan(&x)
	chopDigits(x)
}
