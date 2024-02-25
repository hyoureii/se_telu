package main

import (
	"fmt"
)

func main() {

	var s1, s2, s3 float64
	fmt.Println("Masukkan masing masing berat belerang")
	fmt.Scan(&s1, &s2, &s3)
	fmt.Printf("Berat belerang : %.2f kg atau %.0f gram", (s1+s2+s3)/1000, s1+s2+s3)
}
