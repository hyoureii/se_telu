package main

import (
	"fmt"
)

func main() {
	var startHour, startMinute, endHour, endMinute uint16
	fmt.Printf("Enter start & end hour\n")
	fmt.Scan(&startHour, &startMinute, &endHour, &endMinute)
	if startHour < 7 {
		startHour += 12
	}
	if endHour < 7 {
		endHour += 12
	}
	if startHour >= 18 || endHour < startHour || startHour == 17 && endMinute >= 1 {
		fmt.Printf("Only opens at 7 AM - 5 PM")
	} else {
		duration := (endHour-startHour)*60 - startMinute + endMinute
		fmt.Printf("Parking Duration : %d Hour(s) %d Minute(s)", duration/60, duration%60)
	}
}
