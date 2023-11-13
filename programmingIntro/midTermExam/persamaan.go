package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Scan(&x)
	fmt.Printf("%.2f", 6*x*x-7*x-5)
}
