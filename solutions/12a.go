package solutions

import (
	"fmt"
	"strings"
	"unicode"
)

type route struct {
	forward  map[string]int
	backward map[int]string
	l        int
}

func F12a(input string) int {
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
	r := route{make(map[string]int), make(map[int]string), 0}
	result := walk(&r, "start", &paths)
	return result
}

func walk(r *route, s string, paths *map[string][]string) int {
	if s == "end" {
		return 1
	}
	result := 0
	if unicode.IsLower(rune(s[0])) {
		if _, ok := r.forward[s]; ok {
			return 0
		}
	}
	r.forward[s] = r.l
	r.backward[r.l] = s
	r.l += 1

	for i := range (*paths)[s] {
		result += walk(r, (*paths)[s][i], paths)
	}
	r.l -= 1
	delete(r.backward, r.l)
	delete(r.forward, s)
	return result
}
