package main

import (
	"fmt"
)

func main() {
	var amount, balance int
	for {
		fmt.Scan(&amount)
		if amount == 0 {
			break
		}
		balance += amount
	}
	fmt.Printf("Balance = %d", balance)
}
