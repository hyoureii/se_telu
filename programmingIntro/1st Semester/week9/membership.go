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
	if amount >= 100000 {
		discount = true
	}
	if wantCard == true && discount == true {
		card = true
	}
	if card == true && amount >= 200000 {
		cashback = true
	}
	fmt.Printf("Card? %t\nDiscount? %t\nCashback? %t", card, discount, cashback)
}
