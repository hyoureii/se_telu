package main

import (
	"fmt"
)

func main() {
	var char string
	fmt.Printf("Enter a character\n")
	fmt.Scan(&char)
	if 48 <= char[0] && char[0] <= 57 || 65 <= char[0] && char[0] <= 90 || 97 <= char[0] && char[0] <= 122 {
		fmt.Printf("%t", true)
	} else {
		fmt.Printf("%t", false)
	}
}
