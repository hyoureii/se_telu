package main

import (
	"fmt"
)

func main() {

	var x float64
	fmt.Println("Masukkan nilai x")
	fmt.Scan(&x)
	fmt.Printf("y : %.2f", (3*x-5)*(2*x+1))
}
