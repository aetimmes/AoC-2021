package solutions

import (
	"fmt"
	"strings"
	"unicode"
)

type routeB struct {
	seen          map[string]int
	canDoubleBack bool
}

func F12b(input string) int {
	paths := make(map[string][]string)
	lines := strings.Split(input, "\n")
	for i := range lines {
		if lines[i] != "" {
			fmt.Println("parsing:", lines[i])
			tokens := strings.Split(lines[i], "-")
			if _, ok := paths[tokens[0]]; !ok {
				paths[tokens[0]] = make([]string, 0)
			}
			if _, ok := paths[tokens[1]]; !ok {
				paths[tokens[1]] = make([]string, 0)
			}
			paths[tokens[0]] = append(paths[tokens[0]], tokens[1])
			paths[tokens[1]] = append(paths[tokens[1]], tokens[0])
		}
	}
	r := routeB{make(map[string]int), true}

	result := walkB(&r, "start", &paths)
	return result
}

func walkB(r *routeB, s string, paths *map[string][]string) int {
	if len(r.seen) > 1 && s == "start" {
		return 0
	}
	if s == "end" {
		return 1
	}
	result := 0
	if !canVisit(r, s) {
		return 0
	}
	r.seen[s]++
	set := unicode.IsLower(rune(s[0])) && r.seen[s] == 2
	if set {
		r.canDoubleBack = false
	}
	for i := range (*paths)[s] {
		result += walkB(r, (*paths)[s][i], paths)
	}
	if set {
		r.canDoubleBack = true
	}
	r.seen[s]--
	return result
}

func canVisit(r *routeB, s string) bool {
	if !unicode.IsLower(rune(s[0])) {
		return true
	}
	if r.canDoubleBack && r.seen[s] <= 1 {
		return true
	} else {
		return r.seen[s] == 0
	}
}
