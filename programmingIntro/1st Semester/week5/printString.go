package main

import (
	"fmt"
)

func main() {
	var amount uint
	var text string
	fmt.Println("Enter the amount to print and the String")
	fmt.Scan(&amount, &text)
	for i := 0; uint(i) < amount; i++ {
		fmt.Println(text)
	}
}
