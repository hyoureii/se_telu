package main

import (
	"fmt"
)

func main() {

	var x float64
	fmt.Println("Masukkan nilai x")
	fmt.Scan(&x)
	fmt.Printf("Hasil : %.2f", calc(x))
}

func calc(x float64) float64 {

	return (x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4)
}
