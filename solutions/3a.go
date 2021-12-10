package solutions

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func F3a(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Scan()
	gamma := make([]int, len(scanner.Text()))
	aggregate(scanner.Text(), gamma)
	for scanner.Scan() {
		aggregate(scanner.Text(), gamma)
	}

	g, e := 0, 0

	for i := range gamma {
		if gamma[len(gamma)-i-1] > 0 {
			g += int(math.Pow(2, float64(i)))
		} else {
			e += int(math.Pow(2, float64(i)))
		}
	}
	return (g * e)
}

func aggregate(s string, gamma []int) {
	for i, c := range s {
		c, _ := strconv.Atoi(string(c))
		gamma[i] += c*2 - 1
	}
}
