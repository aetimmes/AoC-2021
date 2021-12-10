package solutions

import (
	"bufio"
	"container/ring"
	"strconv"
	"strings"
)

func F1b(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
	r := ring.New(3)

	for i := 0; i < 3; i++ {
		scanner.Scan()
		r.Value, _ = strconv.Atoi(scanner.Text())
		r = r.Next()
	}

	result := 0
	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())
		if current > int(r.Value.(int)) {
			result += 1
		}
		r.Value = current
		r = r.Next()
	}
	return result
}
