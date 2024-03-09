package main

import (
	"fmt"
)

func won(goal, loss int, wins *int) {
	if goal > loss {
		*wins++
	}
}

func draw(goal, loss int, draws *int) {
	if goal == loss {
		*draws++
	}
}

func lose(goal, loss int, defeats *int) {
	if goal < loss {
		*defeats++
	}
}

func goalsCount(goal, loss int, goals, losses, diff *int) {
	*goals += goal
	*losses += loss
	*diff = *goals - *losses
}

func pointCount(wins, draws, points *int) {
	*points = *wins*3 + *draws
}

func main() {
	var matches, goal, loss, wins, draws, defeats, goals, losses, diff, points int
	fmt.Scan(&matches)
	for i := 0; i < matches; i++ {
		fmt.Scan(&goal, &loss)
		won(goal, loss, &wins)
		draw(goal, loss, &draws)
		lose(goal, loss, &defeats)
		goalsCount(goal, loss, &goals, &losses, &diff)
		pointCount(&wins, &draws, &points)
	}
	fmt.Printf("%d %d %d %d %d %d %d %d", matches, wins, draws, defeats, goals, losses, diff, points)
}
