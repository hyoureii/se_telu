package main

import (
	"fmt"
	"strings"
)

func main() {
	var char string
	fmt.Printf("Enter a character\n")
	fmt.Scan(&char)
	char = strings.ToUpper(char)
	if char[0] == 65 || char[0] == 69 || char[0] == 73 || char[0] == 79 || char[0] == 85 {
		fmt.Printf("%t", false)
	} else {
		fmt.Printf("%t", true)
	}
}
