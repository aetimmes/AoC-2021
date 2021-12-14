package solutions

import (
	"fmt"
	"math"
	"strings"
)

func F14b(input string) int {
	mapping := make(map[string][]string)
	var lines []string
	var initial string
	delimiter := " -> "
	lines = strings.Split(input, "\n")
	initial, lines = lines[0], lines[2:]

	for i := range lines {
		if strings.Contains(lines[i], delimiter) {
			tokens := strings.Split(lines[i], delimiter)
			mapping[tokens[0]] = []string{
				string(tokens[0][0]) + tokens[1],
				tokens[1] + string(tokens[0][1]),
			}
		}
	}
	state := make(map[string]int)
	counts := make(map[byte]int)
	counts[initial[0]] = 1
	for i := 1; i < len(initial); i++ {
		current := string([]byte{initial[i-1], initial[i]})
		state[current]++
		counts[initial[i]]++
	}
	for i := 0; i < 40; i++ {
		state = step(state, &counts, mapping)
	}
	fmt.Println(state)
	fmt.Println(counts)
	max, min := math.MinInt, math.MaxInt
	for _, v := range counts {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return max - min
}
