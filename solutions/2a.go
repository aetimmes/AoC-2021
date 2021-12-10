package solutions

import (
	"bufio"
	"strconv"
	"strings"
)

func F2a(input string) int {
	x, y := 0, 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		current := scanner.Text()
		tokens := strings.Fields(current)
		magnitude, _ := strconv.Atoi(tokens[1])
		switch dir := tokens[0]; dir {
		case "forward":
			x += magnitude
		case "up":
			y -= magnitude
		case "down":
			y += magnitude
		}

	}
	return x * y
}
