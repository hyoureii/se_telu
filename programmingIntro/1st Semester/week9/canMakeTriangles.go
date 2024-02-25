package main

import (
	"fmt"
)

func main() {
	var x, y, z uint
	fmt.Printf("Enter 3 positive integers\n")
	fmt.Scan(&x, &y, &z)
	if checkTriangle(x, y, z) && checkTriangle(x, z, y) && checkTriangle(y, z, x) {
		fmt.Printf("%t", true)
	} else {
		fmt.Printf("%t", false)
	}
}

func checkTriangle(a, b, c uint) bool {
	if a+b > c {
		return true
	} else {
		return false
	}
}
