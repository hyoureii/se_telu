package main

import (
	"fmt"
)

func main() {
	var num1, num2, max, gcd int
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	if num1 > num2 {
		max = num1
	} else {
		max = num2
	}
	x := 1
	for x <= max {
		if num1%x == 0 && num2%x == 0 {
			gcd = x
		}
		x++
	}
	fmt.Printf("%d", gcd)
}
