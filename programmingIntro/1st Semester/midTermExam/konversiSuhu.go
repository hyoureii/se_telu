package main

import (
	"fmt"
)

func main() {
	var suhu float64
	fmt.Scan(&suhu)
	fmt.Printf("%.2f", (suhu*9)/5+32)
}
