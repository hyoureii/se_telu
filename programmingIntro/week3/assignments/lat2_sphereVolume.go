package main

import (
	"fmt"
)

func main() {

	var r float64
	fmt.Println("Masukkan jari jari")
	fmt.Scan(&r)
	fmt.Printf("Volume : %.2f", (4*3.14*(r*r*r))/3)
}
