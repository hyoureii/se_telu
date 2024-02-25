package main

import (
	"fmt"
)

func main() {
	var enemies, kenshin_power, kenshin_speed, enemy_power, enemy_speed, defeated int
	fmt.Scan(&enemies, &kenshin_power, &kenshin_speed)
	for enemies > 0 {
		fmt.Scan(&enemy_power, &enemy_speed)
		if kenshin_power > enemy_power {
			defeated++
			kenshin_power -= 6
		} else if kenshin_speed > enemy_speed {
			defeated++
			kenshin_speed -= 6
		}
		enemies--
	}
	fmt.Printf("%d %d %d", defeated, kenshin_power, kenshin_speed)
}
