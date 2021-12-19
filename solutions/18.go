package solutions

import (
	"log"
	"strings"
)

type snailfishElem struct {
	parent *snailfishNum
	value  *int
	ptr    *snailfishNum
}

type snailfishNum struct {
	parent *snailfishElem
	left   *snailfishElem
	right  *snailfishElem
}

func F18a(input string) int {
	lines := strings.Split(input, "\n")
	sum, _ := parseSnailfishNum(lines[0], nil)
	for i := 1; i < len(lines); i++ {
		if lines[i] != "" {
			current, _ := parseSnailfishNum(lines[i], nil)
			sum = addSnailfishNums(sum, current)
		}
	}
	return getSFNMagnitude(*sum)
}

func parseSnailfishNum(s string, parent *snailfishElem) (*snailfishNum, string) {
	result := snailfishNum{parent, nil, nil}
	// The first character in a SnailfishNum is always a [
	if s[0] != '[' {
		log.Fatalf("parseSnailfishNum failed: Expected '[', got '%b'", s[0])
	}
	s = s[1:]
	result.left, s = parseSnailfishElem(s, &result)
	// The middle character in a SnailfishNum is always a ,
	if s[0] != ',' {
		log.Fatalf("parseSnailfishNum failed: Expected ',', got '%b'", s[0])
	}
	s = s[1:]
	result.right, s = parseSnailfishElem(s, &result)
	// The final character in a snailfishNum is always a ]
	if s[0] != ',' {
		log.Fatalf("parseSnailfishNum failed: Expected ,, got '%b'", s[0])
	}
	s = s[1:]

	return &result, s
}

func parseSnailfishElem(s string, parent *snailfishNum) (*snailfishElem, string) {
	result := snailfishElem{parent, nil, nil}
	if s[0] == '[' {
		result.ptr, s = parseSnailfishNum(s, &result)
	} else {
		result.value, s = parseValue(s)
	}
	return &result, s
}

func parseValue(s string) (*int, string) {
	result := int(s[0] - '0')
	return &result, s[1:]
}

func addSnailfishNums(left, right *snailfishNum) *snailfishNum {
	result := snailfishNum{nil, nil, nil}
	l := snailfishElem{&result, nil, left}
	left.parent = &l
	result.left = &l
	r := snailfishElem{&result, nil, right}
	right.parent = &r
	result.right = &r
	return reduceSnailfishNum(&result)
}

func reduceSnailfishNum(num *snailfishNum) *snailfishNum {
	var exploded, split bool = false, false
	for true {
		num, exploded = explodeSnailfishNum(num, 0)
		if !exploded {
			num, split = splitSnailfishNum(num)
			if !split {
				break
			}
		}
	}
	return num
}

func splitSnailfishNum(num *snailfishNum) (*snailfishNum, bool) {
	
	found := false
	return num, found
}

func explodeSnailfishNum(num *snailfishNum, level int) (*snailfishNum, bool) {
	found := false
	return num, found
}

func getSFNMagnitude(sum snailfishNum) int {
	panic("unimplemented")
}
