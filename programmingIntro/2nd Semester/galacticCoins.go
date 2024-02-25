package main

import (
	"fmt"
)

func main() {
	var lumin int
	fmt.Scan(&lumin)
	quantum := lumin / 1200
	lumin -= 1200 * quantum
	galactum := lumin / 120
	lumin -= 120 * galactum
	nebula := lumin / 40
	lumin -= 40 * nebula
	stellaris := lumin / 20
	lumin -= 20 * stellaris
	fmt.Printf("%d %d %d %d %d", quantum, galactum, nebula, stellaris, lumin)
}
