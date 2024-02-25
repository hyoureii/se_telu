package main

import (
	"fmt"
)

func fibLoop(num uint) {
	var first, second uint = 0, 1
	fmt.Printf("%d %d ", first, second)
	for i := 2; uint(i) <= num; i++ {
		new := first + second
		fmt.Printf("%d ", new)
		first = second
		second = new
	}
}

func main() {
	var num uint
	fmt.Printf("Enter integer\n")
	fmt.Scan(&num)
	if num <= 1 {
		fmt.Printf("Must be larger than 1")
	} else {
		fibLoop(num)
	}
}
