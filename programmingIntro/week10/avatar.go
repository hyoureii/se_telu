package main

import (
	"fmt"
)

func main() {
	var people, adult, child int
	fmt.Printf("Enter amount of people\n")
	fmt.Scan(&people)
	for people > 0 && adult < 3 {
		people -= 5
		adult++
	}
	fmt.Printf("%d Adult", adult)
	if people > 0 {
		for people > 0 && child < 5 {
			people -= 2
			child++
		}
		fmt.Printf(", %d Child", child)
	}
	if people > 0 {
		fmt.Printf(", %d People cant set off", people)
	}
}
