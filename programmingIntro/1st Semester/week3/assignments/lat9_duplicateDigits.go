package main

import (
	"fmt"
)

func main() {

	var x uint16
	fmt.Println("Masukkan bilangan berdigit 2")
	fmt.Scan(&x)
	fmt.Printf("Hasil Penggandaan : %d", (x/10)*1100+(x%10)*11)
}
