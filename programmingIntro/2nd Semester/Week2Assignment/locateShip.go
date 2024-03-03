package main

import (
	"fmt"
)

func main() {
	var xpos, ypos float64
	fmt.Scan(&xpos, &ypos)
	if xpos >= 0 {
		if ypos >= 0 {
			fmt.Printf("Kuadran I")
		} else {
			fmt.Printf("Kuadran IV")
		}
	} else {
		if ypos >= 0 {
			fmt.Printf("Kuadran II")
		} else {
			fmt.Printf("Kuadran III")
		}
	}
}
