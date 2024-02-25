package main

import (
	"fmt"
)

func findFactor(num uint) {
	for i := 1; uint(i) <= num; i++ {
		if num%uint(i) == 0 {
			fmt.Printf("%d %t\n", i, true)
		} else {
			fmt.Printf("%d %t\n", i, false)
		}
	}
}

func main() {
	var x uint
	fmt.Printf("Enter a positive integer\n")
	fmt.Scan(&x)
	findFactor(x)
}
