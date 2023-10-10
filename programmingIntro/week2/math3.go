package main

import (
	"fmt"
)

func main() {

	var x uint16
	fmt.Println("Masukkan nilai x")
	fmt.Scan(&x)
	fmt.Printf("Hasil : %d", calc(x))
}

func calc(x uint16) [3]uint16 {

	var res [3]uint16
	res[0] = x / 100
	res[1] = (x - 100*res[0]) / 10
	res[2] = (x - 100*res[0] - 10*res[1])
	return res
}
