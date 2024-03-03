package main

import (
	"fmt"
	"math"
)

func DegreeToRadian(deg float64) float64 {
	return deg * math.Pi / 180
}

func panjangX(length, angle float64) float64 {
	return length * math.Cos(DegreeToRadian(angle))
}

func panjangY(length, angle float64) float64 {
	return length * math.Sin(DegreeToRadian(angle))
}

func main() {
	var length, angle float64
	fmt.Scan(&length, &angle)
	fmt.Printf("%.2f %.2f", panjangX(length, angle), panjangY(length, angle))
}
