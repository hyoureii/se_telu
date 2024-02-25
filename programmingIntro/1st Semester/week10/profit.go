package main

import (
	"fmt"
)

func main() {
	var before, after float64
	fmt.Printf("Enter profits\n")
	fmt.Scan(&before, &after)
	if after == before {
		fmt.Printf("Same")
	} else if before > after {
		fmt.Printf("Decreased by %.1f", before-after)
	} else {
		fmt.Printf("Increased by %.1f", after-before)
	}
}
