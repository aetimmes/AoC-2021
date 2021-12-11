package solutions

import (
	"log"
	"strings"

	"github.com/aetimmes/go-set/set"
)

const steps11a = 100

var semiCardinalDirs = map[string]point{
	"N":  {0, 1},
	"NE": {1, 1},
	"E":  {1, 0},
	"SE": {1, -1},
	"S":  {0, -1},
	"SW": {-1, -1},
	"W":  {-1, 0},
	"NW": {-1, 1},
}

func F11a(input string) int {
	result := 0
	octopi, r, c := parseGrid(input)
	for i := 0; i < steps11a; i++ {
		result += processFlashes(&octopi, r, c)
	}
	return result
}

func processFlashes(octopi *[][]int, r int, c int) int {
	flashed := set.NewSet[point]()
	for y := 0; y < r; y++ {
		for x := 0; x < c; x++ {
			incrementOctopus(octopi, &flashed, r, c, point{x, y})
		}
	}
	return set.Size(&flashed)
}

func parseGrid(input string) ([][]int, int, int) {
	lines := strings.Split(input, "\n")
	result := make([][]int, 0)
	for i := range lines {
		if lines[i] != "\n" && lines[i] != "" {
			result = append(result, make([]int, 0))
			for j := range lines[i] {
				result[i] = append(result[i], int(lines[i][j]-'0'))
			}
		}
	}
	sanityCheckGrid(result)
	return result, len(result), len(result[0])
}

func sanityCheckGrid(result [][]int) {
	l := len(result[0])
	for i := range result {
		if len(result[i]) != l {
			log.Fatalf("row %d mismatched, expected %d length but got %d", i, l, len(result[i]))
		}
	}
}

func incrementOctopus(octopi *[][]int, flashed *set.Set[point], r, c int, p point) {
	if !checkBounds(p, r, c) || set.Contains(flashed, p) {
		return
	}
	if (*octopi)[p.y][p.x] == 9 {
		set.Add(flashed, p)
		(*octopi)[p.y][p.x] = 0
		for _, d := range semiCardinalDirs {
			incrementOctopus(octopi, flashed, r, c, point{p.x + d.x, p.y + d.y})
		}
	} else {
		(*octopi)[p.y][p.x] += 1
	}
}
