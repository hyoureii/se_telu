package main

import (
	"fmt"
)

const MAX_SIZE int = 11

type player struct {
	name, jerseyNumber, pos string
	height                  int
}

type playerList [MAX_SIZE]player

func readStat(arr *playerList, size *int) {
	var name, jerseyNumber, pos string
	var height int
	for *size < MAX_SIZE {
		fmt.Scan(&name)
		if name == "none" {
			break
		} else {
			fmt.Scan(&jerseyNumber, &pos, &height)
			arr[*size].name = name
			arr[*size].jerseyNumber = jerseyNumber
			arr[*size].pos = pos
			arr[*size].height = height
			*size++
		}
	}
}

func printPlayers(arr playerList, size int) {
	fmt.Printf("Data Pemain:")
	for i := 0; i < size; i++ {
		fmt.Printf("\n%s %s %s %d", arr[i].name, arr[i].jerseyNumber, arr[i].pos, arr[i].height)
	}
	fmt.Printf("\n")
}

func maxHeightPlayer(arr playerList, size int) {
	maxHeight := arr[0].height
	indexMaxHeight := 0
	for i := 1; i < size; i++ {
		if arr[i].height > maxHeight {
			maxHeight = arr[i].height
			indexMaxHeight = i
		}
	}
	fmt.Printf("Pemain dengan badan tertinggi: %s\n", arr[indexMaxHeight].name)
}

func minHeightPlayer(arr playerList, size int) {
	minHeight := arr[0].height
	indexMinHeight := 0
	for i := 1; i < size; i++ {
		if arr[i].height < minHeight {
			minHeight = arr[i].height
			indexMinHeight = i
		}
	}
	fmt.Printf("Pemain dengan badan terendah: %s", arr[indexMinHeight].name)
}

func main() {
	var playerStat playerList
	players := 0
	readStat(&playerStat, &players)
	printPlayers(playerStat, players)
	maxHeightPlayer(playerStat, players)
	minHeightPlayer(playerStat, players)
}
