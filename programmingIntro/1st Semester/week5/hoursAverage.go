package main

import (
	"fmt"
)

func main() {
	var days int
	var avg float64
	fmt.Println("Enter how much days and how much hours each day")
	fmt.Scan(&days)
	hours := make([]int, days)
	for i := 0; i < days; i++ {
		fmt.Scan(&hours[i])
		avg += float64(hours[i])
	}
	fmt.Printf("Average: %.2f", avg/float64(days))
}
