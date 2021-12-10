package solutions

import (
	"bufio"
	"strings"
)

func F8a(input string) int {
	result := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		_, r := parseLine8a(scanner.Text())
		for i := range r {
			switch len(r[i]) {
			case 2, 3, 4, 7:
				result += 1
			default:
			}
		}
	}
	return result
}

func parseLine8a(s string) ([]string, []string) {
	temp := strings.Split(s, "|")
	l := strings.Fields(temp[0])
	r := strings.Fields(temp[1])
	return l, r
}
