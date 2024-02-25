package main

import (
	"fmt"
)

func main() {
	var x uint
	fmt.Scan(&x)
	fmt.Printf("%d", (x/10)*1100+(x%10)*11)
}
