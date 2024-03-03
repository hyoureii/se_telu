package main

import (
	"fmt"
)

func LowerToUppercase(char byte) byte {
	return char - 32
}

func main() {
	var char byte
	fmt.Scanf("%c", &char)
	fmt.Printf("%c", LowerToUppercase(char))
}
