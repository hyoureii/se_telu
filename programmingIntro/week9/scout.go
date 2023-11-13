package main

import (
	"fmt"
)

func main() {
	var members uint
	var item1, item2, item3, item4, item5 bool
	ready := true
	fmt.Scan(&members)
	for i := 0; i < int(members); i++ {
		fmt.Scan(&item1, &item2, &item3, &item4, &item5)
		if item1 && item2 && item3 && item4 && item5 {
			continue
		} else {
			ready = false
			break
		}
	}
	fmt.Printf("%t", ready)
}
