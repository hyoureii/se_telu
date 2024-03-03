package main

import (
	"fmt"
)

func main() {
	var lumin int
	fmt.Scan(&lumin)
	quantum, galactum, nebula, stellaris := 0, 0, 0, 0
	for lumin >= 20 {
		if lumin >= 1200 {
			quantum++
			lumin -= 1200
		} else if lumin >= 120 {
			galactum++
			lumin -= 120
		} else if lumin >= 40 {
			nebula++
			lumin -= 40
		} else if lumin >= 20 {
			stellaris++
			lumin -= 20
		}
	}
	fmt.Printf("%d %d %d %d %d", quantum, galactum, nebula, stellaris, lumin)
}
