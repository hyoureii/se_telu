package main

import (
	"fmt"
)

func main() {
	var cap, vol, filled uint
	filled = 0
	fmt.Printf("Enter capacity, then volume of each bucket to fill\n")
	fmt.Scan(&cap)
	for true {
		fmt.Scan(&vol)
		filled += vol
		if filled >= cap {
			fmt.Printf("%t\n", true)
			break
		} else {
			fmt.Printf("%t\n", false)
		}
	}
}
