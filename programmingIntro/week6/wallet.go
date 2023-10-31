package main

import (
	"fmt"
)

func main() {
	var amount, balance int
	fmt.Printf("Enter transactions, will end if input is 0\n")
	for {
		fmt.Scan(&amount)
		if amount == 0 {
			break
		}
		balance += amount
	}
	fmt.Printf("Balance = %d", balance)
}
