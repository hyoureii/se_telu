package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)
	nums := strconv.Itoa(num)
	if len(nums) <= 2 {
		fmt.Printf("%t", true)
		fmt.Println("less or equal than 2")
	} else {
		diff := int(nums[1]) - int(nums[0])
		consecutive := true
		for digit := 2; digit < len(nums); digit++ {
			if int(nums[digit])-int(nums[digit-1]) != diff {
				fmt.Printf("%t", false)
				fmt.Println("not same")
				break
			}
		}
		if consecutive == true {
			fmt.Printf("%t", true)
			fmt.Println("same")
		}
	}
}
