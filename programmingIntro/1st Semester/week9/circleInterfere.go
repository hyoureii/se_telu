package main

import (
	"fmt"
)

func main() {
	var r1, r2, distance uint
	fmt.Printf("Enter radius1, radius2, and distance\n")
	fmt.Scan(&r1, &r2, &distance)
	fmt.Printf("%t", r1+r2 > distance)
}
