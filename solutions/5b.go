package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func F5b(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	seen := make(map[point]int)
	for scanner.Scan() {
		start, finish := parseLine5(scanner.Text())
		x_dir := 0
		if start.x < finish.x {
			x_dir = 1
		} else if start.x > finish.x {
			x_dir = -1
		}
		y_dir := 0
		if start.y < finish.y {
			y_dir = 1
		} else if start.y > finish.y {
			y_dir = -1
		}
		p := newPoint(start.x, start.y)
		for ; p.x != finish.x || p.y != finish.y; p.x, p.y = p.x+x_dir, p.y+y_dir {
			seen[p] += 1
		}
		seen[p] += 1
	}
	for _, v := range seen {
		if v > 1 {
			result += 1
		}
	}
	fmt.Println(result)
}

func parseLiner5(s string) (point, point) {
	tokens := strings.Split(s, " -> ")
	return newPointFromString(tokens[0]), newPointFromString(tokens[1])
}
