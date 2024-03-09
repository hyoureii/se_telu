package main

import (
	"fmt"
)

func circle(radius float64, area, circum *float64) {
	*area = 3.14 * radius * radius
	*circum = 2 * 3.14 * radius
}

func square(side float64, area, circum *float64) {
	*area = side * side
	*circum = 4 * side
}

func total(circleArea, squareArea, circleCircum, squareCircum float64, totalArea, totalCircum *float64) {
	*totalArea = circleArea + squareArea
	*totalCircum = circleCircum + squareCircum
}

func main() {
	var r, s, ll, lp, kl, kp, tl, tk float64
	tagsPrinted := false
	for true {
		fmt.Scan(&r, &s)
		if r == 0 || s == 0 {
			break
		}
		if tagsPrinted == false {
			fmt.Printf("%7s%7s%7s%7s%7s%7s%7s%7s\n", "R", "S", "LL", "LP", "KL", "KP", "TL", "TP")
			tagsPrinted = true
		}
		circle(r, &ll, &kl)
		square(s, &lp, &kp)
		total(ll, lp, kl, kp, &tl, &tk)
		fmt.Printf("%7.2f%7.2f%7.2f%7.2f%7.2f%7.2f%7.2f%7.2f\n", r, s, ll, lp, kl, kp, tl, tk)
	}
}
