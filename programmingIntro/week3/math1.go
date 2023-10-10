package main

import (
	"fmt"
)

func main() {

	var r float64
	fmt.Println("Please input the radius")
	fmt.Scan(&r)
	fmt.Printf("Result : %.2f", sphereSurfArea(r))
}

func sphereSurfArea(r float64) float64 {

	return (4 * 22 * r * r) / 7
}
