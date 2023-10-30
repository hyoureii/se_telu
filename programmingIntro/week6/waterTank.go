package main

import (
	"fmt"
)

func main() {
	var cap, vol, filled uint
	filled = 0
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
