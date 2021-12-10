package solutions

import (
	"bufio"
	"strings"
)

var scores = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var bracePairs = map[byte]byte{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

func F10a(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	result := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			result += scoreLine(line)
		}
	}
	return result
}

func scoreLine(s string) int {
	q := make([]byte, 0, 100)
	for i := range s {
		_, ok := scores[s[i]]
		if ok {
			if q[len(q)-1] == bracePairs[s[i]] {
				q = q[:len(q)-1]
			} else {
				return scores[s[i]]
			}
		} else {
			q = append(q, s[i])
		}
	}
	return 0
}
