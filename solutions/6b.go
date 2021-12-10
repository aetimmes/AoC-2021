package solutions

import (
	"bufio"
	"strings"
)

func F6b(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	initial_fish := AStoi(strings.Split(scanner.Text(), ","))

	fish := make([]int, 9)
	for i := range initial_fish {
		fish[initial_fish[i]] += 1
	}

	for i := 0; i < 256; i++ {
		fish = []int{
			fish[1],
			fish[2],
			fish[3],
			fish[4],
			fish[5],
			fish[6],
			fish[7] + fish[0],
			fish[8],
			fish[0],
		}
	}

	result := 0
	for i := range fish {
		result += fish[i]
	}
	return result
}
