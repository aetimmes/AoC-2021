package solutions

import (
	"bufio"
	"strconv"
	"strings"
)

func F2b(input string) int {
	x, y, aim := 0, 0, 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		current := scanner.Text()
		tokens := strings.Fields(current)
		magnitude, _ := strconv.Atoi(tokens[1])
		switch dir := tokens[0]; dir {
		case "forward":
			x += magnitude
			y += magnitude * aim
		case "up":
			aim -= magnitude
		case "down":
			aim += magnitude
		}

	}
	return x * y
}
