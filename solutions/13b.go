package solutions

import (
	"strconv"
	"strings"
)

func F13b(input string) int {
	halves := strings.Split(input, "\n\n")
	pl := strings.Split(halves[0], "\n")
	points := make(map[point]int)
	for i := range pl {
		tokens := strings.Split(pl[i], ",")
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		points[point{x, y}] = 1

	}
	printPoints(points)
	folds := strings.Split(halves[1], "\n")
	for i := range folds {
		points = fold(folds[i], points)
	}
	printPoints(points)
	return len(points)
}
