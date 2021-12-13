package solutions

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func F13a(input string) int {
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
	for i := 0; i < 1; i++ {
		points = fold(folds[i], points)
	}
	printPoints(points)
	return len(points)
}

func fold(s string, points map[point]int) map[point]int {
	result := make(map[point]int)
	fmt.Println(s)
	f := strings.Split(strings.Split(s, " ")[2], "=")
	dir := f[0]
	loc, _ := strconv.Atoi(f[1])
	fmt.Printf("folding across %s = %d\n", dir, loc)
	if dir == "y" {
		for i := range points {
			if i.y > loc {
				y := loc - (i.y - loc)
				fmt.Printf("folding {%d, %d} to {%d, %d}\n", i.x, i.y, i.x, y)
				result[point{i.x, y}] = 1
			} else {
				result[i] = 1
			}
		}
	} else if dir == "x" {
		for i := range points {
			if i.x > loc {
				x := loc - (i.x - loc)
				fmt.Printf("folding {%d, %d} to {%d, %d}\n", i.x, i.y, x, i.y)
				result[point{x, i.y}] = 1
			} else {
				result[i] = 1
			}
		}
	} else {
		log.Fatalf("unknown fold direction %s", dir)
	}
	return result
}

func printPoints(points map[point]int) {
	for y := 0; y < 7; y++ {
		for x := 0; x < 41; x++ {
			if points[point{x, y}] == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
