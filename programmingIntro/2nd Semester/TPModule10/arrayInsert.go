package main

import (
	"fmt"
)

const MAX_SIZE int = 20

type intArray [MAX_SIZE]int

func readValue(arr *intArray, size *int) {
	var inp int
	for *size < MAX_SIZE {
		fmt.Scan(&inp)
		if inp < 1 {
			break
		} else {
			arr[*size] = inp
			*size++
		}
	}
}

func printArray(arr intArray, size int) {
	fmt.Printf("Elemen array : ")
	for i := 0; i < size; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Printf("\n")
}

func findMax(arr intArray, size int) int {
	currentMax := arr[0]
	for i := 1; i < size; i++ {
		if arr[i] > currentMax {
			currentMax = arr[i]
		}
	}
	return currentMax
}

func findMin(arr intArray, size int) int {
	currentMin := arr[0]
	for i := 1; i < size; i++ {
		if arr[i] < currentMin {
			currentMin = arr[i]
		}
	}
	return currentMin
}

func printInfo(arr intArray, size int) {
	fmt.Printf("Nilai maksimum : %d\nNilai minimum : %d\nBanyak elemen : %d", findMax(arr, size), findMin(arr, size), size)
}

func main() {
	var arr intArray
	maxSize := 0
	readValue(&arr, &maxSize)
	printArray(arr, maxSize)
	printInfo(arr, maxSize)
}
