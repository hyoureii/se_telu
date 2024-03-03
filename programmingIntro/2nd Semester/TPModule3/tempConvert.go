package main

import (
	"fmt"
)

func reamur(celc float64) float64 {
	return celc * 4 / 5
}

func fahrenheit(celc float64) float64 {
	return (celc * 9 / 5) + 32
}

func kelvin(celc float64) float64 {
	return celc + 273
}

func main() {
	var start, end, interval float64
	fmt.Scan(&start, &end, &interval)
	fmt.Printf("%10s %10s %10s %10s\n", "Celcius", "Reamur", "Fahrenheit", "Kelvin")
	for start <= end {
		fmt.Printf("%10.2f %10.2f %10.2f %10.2f\n", start, reamur(start), fahrenheit(start), kelvin(start))
		start += interval
	}
}
