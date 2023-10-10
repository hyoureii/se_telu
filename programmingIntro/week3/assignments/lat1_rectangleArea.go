package main

import (
	"fmt"
)

func main() {

	var l, w float64
	fmt.Println("Masukkan panjang dan lebar")
	fmt.Scan(&l, &w)
	fmt.Printf("Luas : %.2f", l*w)
}
