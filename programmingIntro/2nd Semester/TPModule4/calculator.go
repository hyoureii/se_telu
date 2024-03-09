package main

import (
	"fmt"
)

func menu() {
	fmt.Print("-----------------------\n        M E N U        \n-----------------------\n1. Hitung Penjumlahan\n2. Hitung Perkalian\n3. Hitung Pembagian\n4. Exit\n-----------------------\n")
}

func addition() {
	var num1, num2 int
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan: ")
	fmt.Scan(&num1, &num2)
	fmt.Printf("Hasil penjumlahan: %d\n", num1+num2)
}

func multiplication() {
	var num1, num2 int
	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scan(&num1, &num2)
	fmt.Printf("Hasil perkalian: %d\n", num1*num2)
}

func divide() {
	var num1, num2 float64
	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scan(&num1, &num2)
	fmt.Printf("Hasil pembagian: %.1f\n", num1/num2)
}

func main() {
	var selection int
	for true {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&selection)
		if selection == 4 {
			break
		}
		switch selection {
		case 1:
			addition()
		case 2:
			multiplication()
		case 3:
			divide()
		}
	}
}
