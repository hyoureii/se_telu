package main

import (
	"fmt"
)

func main() {
	var r, t float64
	fmt.Scan(&r, &t)
	fmt.Printf("%.2f", (2*22*r*t)/7+(2*22*r*r)/7)
}
