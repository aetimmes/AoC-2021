package solutions

import (
	"bufio"
	"strconv"
	"strings"
)

func F1a(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	previous, _ := strconv.Atoi(scanner.Text())
	result := 0
	for scanner.Scan() {
		temp := scanner.Text()
		current, _ := strconv.Atoi(temp)
		if current > previous {
			result += 1
		}
		previous = current
	}
	return result
}
