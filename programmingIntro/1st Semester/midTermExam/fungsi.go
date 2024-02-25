package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)
	fmt.Printf("%.2f", math.Sqrt(((x*x*x)+(3*x))*((x*x*x*x)-(3*x*x)+4)))
}
