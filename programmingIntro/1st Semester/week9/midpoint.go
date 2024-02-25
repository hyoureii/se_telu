package main

import (
	"fmt"
)

func main() {
	var begin, mid, end uint
	fmt.Printf("Enter three positive integers\n")
	fmt.Scan(&begin, &mid, &end)
	if begin > mid {
		begin, mid = mid, begin
	}
	if mid > end {
		mid, end = end, mid
	}
	if begin > mid {
		begin, mid = mid, begin
	}
	if mid-begin == end-mid {
		fmt.Printf("%t", true)
	} else {
		fmt.Printf("%t", false)
	}
}
