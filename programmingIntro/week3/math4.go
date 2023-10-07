package main

import (
	"fmt"
)

func main() {

	var source, capacity, charge uint
	fmt.Println("Please enter E0, c and E1")
	fmt.Scan(&source, &capacity, &charge)
	if capacity < charge {
		fmt.Print("E1 cannot be larger than c")
	} else {
		fmt.Printf("Cycle count : %d", chargeCount(source, capacity, charge))
	}
}

func chargeCount(x, y, z uint) uint {

	if z > 0 {
		x -= (y - z)
	}
	return x / y
}
