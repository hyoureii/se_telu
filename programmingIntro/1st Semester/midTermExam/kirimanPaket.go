package main

import (
	"fmt"
)

func main() {
	var p, l, t, b uint
	fmt.Scan(&p, &l, &t, &b)
	fmt.Printf("%t", p*l*t <= 200 && b <= 2000)
}
