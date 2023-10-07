package main

import (
	"fmt"
)

func main() {

	var r float64
	fmt.Println("Please input the radius")
	fmt.Scan(&r)
	fmt.Printf("Result : %f", calc(r))
}

func calc(r float64) float64 {

	return (4 * 22 * r * r) / 7
}
