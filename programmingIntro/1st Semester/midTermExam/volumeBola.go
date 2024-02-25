package main

import (
	"fmt"
)

func main() {
	var r float64
	fmt.Scan(&r)
	fmt.Printf("%.2f", (4*22*r*r*r)/(3*7))
}
