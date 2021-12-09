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

type BasinCheck struct {
	p point
	d string
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
		toCheck := append(make([]point, 0, 100), basins[i].origin)
		checked := set.NewSet[BasinCheck]()
		for len(toCheck) > 0 {
			current, toCheck := toCheck[0], toCheck[1:]
			for d, dp := range dirs {
				candidate := point{current.x + dp.x, current.y + dp.y}
				if set.Contains(&checked, BasinCheck{candidate, d}) {
					continue
				}
				fmt.Println("Checking:", candidate.x, candidate.y, d)
				if !set.Contains(&basins[i].members, candidate) &&
					checkBounds(candidate, r, c) &&
					checkHeight(candidate, heights) &&
					checkPoint(current, d, heights, r, c) {
					set.Add(&basins[i].members, candidate)
					toCheck = append(toCheck, candidate)
				}
				set.Add(&checked, BasinCheck{candidate, d})
			}
		}
		scores = append(scores, set.Size(&basins[i].members))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(scores)))

	fmt.Println(scores[0] * scores[1] * scores[2])
}

func checkHeight(p point, heights [][]int) bool {
	return heights[p.x][p.y] != 9
}
