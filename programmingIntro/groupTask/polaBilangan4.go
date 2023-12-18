package main

import (
	"fmt"
)

func main() {
	var angka int
	fmt.Printf("Masukkan bilangan positif\n")
	fmt.Scan(&angka)
	for vertical := 1; vertical <= angka; vertical++ {
		for horizontal := 1; horizontal <= angka; horizontal++ {
			if horizontal == 1 || horizontal == angka || vertical == 1 || vertical == angka {
				fmt.Printf("%d ", vertical)
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}
