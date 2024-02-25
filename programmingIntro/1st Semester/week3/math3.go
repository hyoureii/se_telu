package main

import (
	"fmt"
)

func main() {

	var x string
	fmt.Println("Please enter a character")
	fmt.Scan(&x)
	fmt.Printf("Capital : %t", isCap(x))
}

func isCap(x string) bool {

	if 64 < x[0] && x[0] < 91 {
		return true
	} else {
		return false
	}
}
