package solutions

import (
	"bufio"
	"sort"
	"strings"

	"github.com/aetimmes/go-set/set"
)

type Basin struct {
	origin  point
	members set.Set[point]
}

func F9b(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
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

	basins := make([]Basin, 0, 1000)

	for x := 0; x < c; x++ {
		for y := 0; y < r; y++ {
			p := point{x, y}
			if isLocalMin(point{x, y}, heights, r, c) {
				basins = append(basins, Basin{p, set.NewSet(p)})
			}
		}
	}

	scores := make([]int, 0, len(basins))

	for i := range basins {
		toCheck := append(make([]point, 0, 100), basins[i].origin)
		for len(toCheck) > 0 {
			current := toCheck[0]
			toCheck = toCheck[1:]
			for d, dp := range cardinalDirs {
				candidate := point{current.x + dp.x, current.y + dp.y}
				if !checkBounds(candidate, r, c) {
					continue
				}
				if canAdd(candidate, current, basins[i], heights, d, r, c) {
					set.Add(&basins[i].members, candidate)
					toCheck = append(toCheck, candidate)
				}
			}
		}
		scores = append(scores, set.Size(&basins[i].members))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	return scores[0] * scores[1] * scores[2]
}

func canAdd(candidate, current point, b Basin, heights [][]int, d string, r, c int) bool {
	return !set.Contains(&b.members, candidate) &&
		checkBounds(candidate, r, c) &&
		checkHeight(candidate, heights) &&
		checkPoint(current, d, heights, r, c)
}

func checkHeight(p point, heights [][]int) bool {
	return heights[p.y][p.x] != 9
}
