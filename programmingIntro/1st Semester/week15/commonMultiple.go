package main

import (
	"fmt"
)

func main() {
	var num1, num2, max int
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	if num1 > num2 {
		max = num1
	} else {
		max = num2
	}
	x := 1
	for x <= max {
		if x%num1 == 0 && x%num2 == 0 {
			fmt.Printf("%d", x)
			break
		}
		x++
	}
}
