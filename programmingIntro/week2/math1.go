package main

import (
	"fmt"
)

func main() {

	var x, y float64
	fmt.Println("Masukkan nilai x dan y")
	fmt.Scan(&x, &y)
	fmt.Printf("Hasil : %.2f", calc(x, y))
}

func calc(x, y float64) float64 {

	return 1/(3*x*x+10) + 10*y + 7
}
