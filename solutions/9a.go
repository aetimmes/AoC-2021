package solutions

import (
	"bufio"
	"strings"
)

var cardinalDirs = map[string]point{
	"N": {0, 1},
	"S": {0, -1},
	"E": {1, 0},
	"W": {-1, 0},
}

func F9a(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	result := 0
	heights := make([][]int, 0, 100)

	for scanner.Scan() {
		if scanner.Text() != "\n" && scanner.Text() != "" {
			row := make([]int, 0, 100)
			current := strings.TrimSpace(scanner.Text())
			for i := range current {
				row = append(row, int(current[i]-'0'))
			}
			heights = append(heights, row)
		}
	}
	r, c := len(heights), len(heights[0])

	for x := 0; x < c; x++ {
		for y := 0; y < r; y++ {
			if isLocalMin(point{x, y}, heights, r, c) {
				result += 1 + heights[y][x]
			}
		}
	}

	return result
}

func checkBounds(p point, r, c int) bool {
	return (0 <= p.x && p.x < c &&
		0 <= p.y && p.y < r)
}

func isLocalMin(p point, heights [][]int, r, c int) bool {
	for d := range cardinalDirs {
		if !checkPoint(p, d, heights, r, c) {
			return false
		}
	}
	return true
}

func checkPoint(p point, d string, heights [][]int, r, c int) bool {
	p2 := point{p.x + cardinalDirs[d].x, p.y + cardinalDirs[d].y}
	return !checkBounds(p2, r, c) || heights[p.y][p.x] < heights[p2.y][p2.x]
}
