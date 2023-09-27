package main

import (
	"fmt"
)

func main() {

	var x, y float64
	fmt.Println("Please input x and y")
	fmt.Scan(&x, &y)
	fmt.Printf("Result : %f", calc(x, y))
}

func calc(x, y float64) float64 {

	return 1/(3*x*x+10) + 10*y + 7
}
