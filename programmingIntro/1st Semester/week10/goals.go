package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Enter 4 integers\n")
	nums := make([]uint, 4)
	fmt.Scan(&nums[0], &nums[1], &nums[2], &nums[3])
	if nums[0] > nums[1] {
		nums[0], nums[1] = nums[1], nums[0]
	}
	if nums[2] > nums[3] {
		nums[2], nums[3] = nums[3], nums[2]
	}
	if nums[0] > nums[2] {
		nums[0], nums[2] = nums[2], nums[0]
	}
	if nums[1] > nums[3] {
		nums[1], nums[3] = nums[3], nums[1]
	}
	fmt.Printf("Max : %d\nMin : %d", nums[0], nums[3])
}
