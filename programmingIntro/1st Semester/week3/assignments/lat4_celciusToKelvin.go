package main

import (
	"fmt"
)

func main() {

	var c float64
	fmt.Println("Masukkan suhu dalam Celcius")
	fmt.Scan(&c)
	fmt.Printf("Kelvin : %.2f", c+273)
}
