package solutions

import (
	"bufio"
	"sort"
	"strings"
)

var acScores = map[byte]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func F10b(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	scores := make([]int, 0, 106)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			score := scoreACLine(line)
			if score > 0 {
				scores = append(scores, score)
			}
		}
	}
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func scoreACLine(s string) int {
	q := make([]byte, 0, 100)
	for i := range s {
		_, ok := scores[s[i]]
		if ok {
			if q[len(q)-1] == bracePairs[s[i]] {
				q = q[:len(q)-1]
			} else {
				return 0
			}
		} else {
			q = append(q, s[i])
		}
	}
	result := 0
	for i := range q {
		result *= 5
		result += acScores[q[len(q)-i-1]]
	}
	return result
}
