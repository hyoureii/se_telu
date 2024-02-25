package main

import (
	"fmt"
)

func main() {

	var l, w float64
	fmt.Println("Masukkan panjang dan lebar")
	fmt.Scan(&l, &w)
	fmt.Printf("Keliling : %.2f", 2*(l+w))
}
