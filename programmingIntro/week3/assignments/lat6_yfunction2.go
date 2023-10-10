package main

import (
	"fmt"
)

func main() {

	var x float64
	fmt.Println("Masukkan nilai x")
	fmt.Scan(&x)
	fmt.Printf("f(%.0f) : %.2f", x, (x*x*x+3*x)/(x*x*x*x-3*x*x+4))
}
