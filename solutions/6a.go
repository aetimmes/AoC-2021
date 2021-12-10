package solutions

import (
	"bufio"
	"strings"
)

func F6a(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	fish := AStoi(strings.Split(scanner.Text(), ","))

	for i := 0; i < 80; i++ {
		for j := range fish {
			if fish[j] == 0 {
				fish[j] = 6
				fish = append(fish, 8)
			} else {
				fish[j] -= 1
			}
		}
	}
	return len(fish)
}
