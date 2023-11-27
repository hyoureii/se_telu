package main

import (
	"fmt"
)

func main() {
	var a, b, c uint
	fmt.Printf("Enter 3 numbers\n")
	fmt.Scan(&a, &b, &c)
	if a == b && b == c {
		fmt.Printf("Segitiga Sama Sisi")
	} else if a == b || a == c || b == c {
		fmt.Printf("Segitiga Sama Kaki")
	} else {
		fmt.Printf("Segitiga Sembarang")
	}
}
