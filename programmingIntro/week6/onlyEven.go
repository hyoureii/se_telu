package main

import (
	"fmt"
)

func main() {
	var num, total uint
	fmt.Println("Enter a number, will end if the number is not even")
	for true {
		fmt.Scan(&num)
		if num%2 == 0 {
			total += num
		} else {
			break
		}
	}
	fmt.Printf("Total even : %d", total)
}
