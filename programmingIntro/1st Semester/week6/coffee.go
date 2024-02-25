package main

import (
	"fmt"
)

func main() {
	var sugar, coffee, sugarreq, coffeereq uint
	fmt.Printf("Enter available sugar and coffee, then how much a cup needs\n")
	fmt.Scan(&sugar, &coffee, &sugarreq, &coffeereq)
	fmt.Printf("%d cups can be made", min(sugar/sugarreq, coffee/coffeereq))
}
