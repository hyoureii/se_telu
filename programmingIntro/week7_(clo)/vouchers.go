package main

import (
	"fmt"
)

func main() {
	var voucher uint16
	fmt.Printf("Enter voucher number\n")
	fmt.Scan(&voucher)
	if voucher > 9999 {
		fmt.Printf("Number must be 4 digits")
	} else {
		firstDigit := voucher / 1000
		secondDigit := (voucher % 1000) / 100
		thirdDigit := (voucher % 100) / 10
		fourthDigit := (voucher % 10)

		if (secondDigit*10+thirdDigit)%2 == 0 {
			fmt.Printf("Discount = %t\n", true)
		} else {
			fmt.Printf("Discount = %t\n", false)
		}
		if fourthDigit != 0 {
			if (firstDigit+thirdDigit)%fourthDigit == 0 {
				fmt.Printf("Cashback = %t", true)
			} else {
				fmt.Printf("Cashback = %t", false)
			}
		} else if fourthDigit == 0 {
			fmt.Printf("Cashback = %t", false)
		}
	}
}
