package main

import (
	"fmt"
)

func main() {
	var current, next float64
	var rising bool
	fmt.Printf("Enter 5 different temps\n")
	stable := true
	fmt.Scan(&current, &next)
	if current < next {
		rising = true
	} else {
		rising = false
	}
	current = next
	for i := 3; i <= 5; i++ {
		fmt.Scan(&next)
		if rising == true && next < current || rising == false && next > current {
			stable = false
		}
		current = next
	}
	if stable {
		if rising {
			fmt.Printf("Increasing steadily")
		} else {
			fmt.Printf("Decreasing steadily")
		}
	} else {
		fmt.Printf("Not stable")
	}
}
