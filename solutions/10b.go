package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var acScores = map[byte]int{
	'(': 1,
	'[': 2,
	'{': 3,
	'<': 4,
}

func F10b(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
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
	fmt.Println(scores[len(scores)/2])
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
