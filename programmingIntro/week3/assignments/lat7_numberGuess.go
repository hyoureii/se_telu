package main

import (
	"fmt"
)

func main() {

	var adik, kakak int8
	fmt.Println("Masukkan bilangan adik dan kakak")
	fmt.Scan(&adik, &kakak)
	if (adik < 0 || adik > 9) || (kakak < 0 || kakak > 9) {
		fmt.Print("Bilangan harus 0 sampai 9")
	} else {
		fmt.Printf("Adik menang : %t", ((adik%2+kakak%2)%2 != 0))
	}
}
