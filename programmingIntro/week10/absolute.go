package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Printf("Enter an integer\n")
	fmt.Scan(&num)
	if num == 0 {
		fmt.Printf("%d", num)
	} else {
		if num < 0 {
			num *= -1
		}
		fmt.Printf("%d", num)
	}
}
