package main

import (
	"fmt"
)

func main() {

	var l, wi, h, we uint16
	fmt.Println("Masukkan panjang, lebar, tinggi, dan berat paket")
	fmt.Scan(&l, &wi, &h, &we)
	fmt.Printf("Boleh dikirim : %t", isEligibleToShip(l, wi, h, we))
}

func isEligibleToShip(l, wi, h, we uint16) bool {
	if (l*wi*h) > 1000 || we > 30000 {
		return false
	} else {
		return true
	}
}
