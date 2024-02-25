package main

import (
	"fmt"
)

func main() {
	var year uint16
	fmt.Printf("Enter year\n")
	fmt.Scan(&year)
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Printf("%t", true)
	} else {
		fmt.Printf("%t", false)
	}
}
