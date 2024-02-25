package main

import (
	"fmt"
)

func main() {
	var amount uint
	var wantCard bool
	card, discount, cashback := false, false, false
	fmt.Printf("Enter payment amount and if you'd like to get a membership card\n")
	fmt.Scan(&amount, &wantCard)
	total := amount
	if amount >= 100000 {
		discount = true
		total -= amount * 1 / 10
	}
	if wantCard == true && discount == true {
		card = true
	}
	if card == true && amount >= 200000 {
		cashback = true
		total -= 75000
	}
	fmt.Printf("Card? %t\nDiscount? %t\nCashback? %t\nFinal Price : Rp. %d", card, discount, cashback, total)
}
