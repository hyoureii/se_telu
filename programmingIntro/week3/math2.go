package main

import (
	"fmt"
)

func main() {

	var x uint16
	fmt.Println("Please enter a 2-digit positive integer")
	fmt.Scan(&x)
	fmt.Printf("Result : %d", dupe(x))
}

func dupe(x uint16) uint16 {

	return (x/10)*1100 + (x%10)*11
}
