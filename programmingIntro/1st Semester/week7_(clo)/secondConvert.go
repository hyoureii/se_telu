package main

import (
	"fmt"
)

func main() {
	var seconds int
	fmt.Printf("Enter seconds\n")
	fmt.Scan(&seconds)
	hours := seconds / 3600
	minute := (seconds % 3600) / 60
	seconds %= 60
	fmt.Printf("%d hours %d minutes %d seconds", hours, minute, seconds)

}
