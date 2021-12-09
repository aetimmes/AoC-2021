package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/aetimmes/go-set/set"
)

type Basin struct {
	origin  point
	members set.Set[point]
}

func F9b(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
		toCheck := make([]point, 0, 100)
		toCheck = append(toCheck, basins[i].origin)
		for len(toCheck) > 0 {
			current := toCheck[0]
			toCheck = toCheck[1:]
			for d, dp := range dirs {
				candidate := point{current.x + dp.x, current.y + dp.y}
				if !checkBounds(candidate, r, c) {
					continue
				}
				if !set.Contains(&basins[i].members, candidate) &&
					checkBounds(candidate, r, c) &&
					checkHeight(candidate, heights) &&
					checkPoint(current, d, heights, r, c) {
					set.Add(&basins[i].members, candidate)
					toCheck = append(toCheck, candidate)
				}
			}
		}
		scores = append(scores, set.Size(&basins[i].members))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	fmt.Println(scores[0] * scores[1] * scores[2])
}

func checkHeight(p point, heights [][]int) bool {
	return heights[p.y][p.x] != 9
}
