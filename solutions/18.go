package solutions

import (
	"fmt"
	"strconv"
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
	return getSFNMagnitude(sum)
}

func F18b(input string) int {
	lines := strings.Split(input, "\n")
	result := 0
	for i := range lines {
		for j := range lines {
			if i != j && lines[i] != "" && lines[j] != "" {
				l, _ := parseSnailfishNum(lines[i], nil)
				r, _ := parseSnailfishNum(lines[j], nil)
				current := getSFNMagnitude(addSnailfishNums(l, r))
				if current > result {
					result = current
				}
			}
		}
	}
	return result
}

func sfnToString(num *snailfishNum) string {
	if num == nil {
		return ""
	}
	return "[" + sfeToString(num.left) + "," + sfeToString(num.right) + "]"
}

func sfeToString(elem *snailfishElem) string {
	if elem.value != nil {
		return strconv.Itoa(*elem.value)
	} else {
		return sfnToString(elem.ptr)
	}
}

func parseSnailfishNum(s string, parent *snailfishElem) (*snailfishNum, string) {
	result := snailfishNum{parent, nil, nil}
	// The first character in a SnailfishNum is always a [
	if s[0] != '[' {
		msg := fmt.Sprintf("parseSnailfishNum failed: Expected '[', got '%c'", s[0])
		panic(msg)
	}
	s = s[1:]
	result.left, s = parseSnailfishElem(s, &result)
	// The middle character in a SnailfishNum is always a ,
	if s[0] != ',' {
		msg := fmt.Sprintf("parseSnailfishNum failed: Expected ',', got '%c'", s[0])
		panic(msg)
	}
	s = s[1:]
	result.right, s = parseSnailfishElem(s, &result)
	// The final character in a snailfishNum is always a ]
	if s[0] != ']' {
		msg := fmt.Sprintf("parseSnailfishNum failed: Expected ']', got '%c'", s[0])
		panic(msg)
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
		exploded = explodeSFN(num, 1)
		if !exploded {
			split = splitSFN(num)
			if !split {
				break
			}
		}
	}
	return num
}

func splitSFN(num *snailfishNum) bool {
	found := splitSnailfishElem(num.left)
	if !found {
		found = splitSnailfishElem(num.right)
	}
	return found
}

func splitSnailfishElem(elem *snailfishElem) bool {
	if elem.value != nil {
		if *elem.value > 9 {
			l := *elem.value / 2
			r := *elem.value/2 + *elem.value%2
			n := snailfishNum{elem, nil, nil}
			left := snailfishElem{&n, &l, nil}
			right := snailfishElem{&n, &r, nil}
			n.left, n.right = &left, &right
			elem.value = nil
			elem.ptr = &n
			return true
		} else {
			return false
		}
	} else {
		return splitSFN(elem.ptr)
	}
}

func explodeSFN(num *snailfishNum, level int) bool {
	if num == nil {
		return false
	}
	if level >= 5 && num.left.value != nil && num.right.value != nil {
		l, r := *num.left.value, *num.right.value
		temp := 0
		addToNextLeft(num, l)
		addToNextRight(num, r)
		num.parent.ptr = nil
		num.parent.value = &temp
		return true
	}
	return explodeSFN(num.left.ptr, level+1) || explodeSFN(num.right.ptr, level+1)
}

func sfnNodeToString(num *snailfishNum) string {
	if num == nil {
		return "Nil"
	}
	result := fmt.Sprintf("Address: %p, ", num)
	if num.parent != nil {
		result += fmt.Sprintf("Parent: %p, ", num.parent)
	}
	if num.left.value != nil {
		result += fmt.Sprintf("Left value: %d", *num.left.value)
	} else {
		result += fmt.Sprintf("Left pointer: %d", *num.left.value)
	}
	return result
}

func addToNextLeft(num *snailfishNum, i int) {
	for isLeftNode(num) {
		num = num.parent.parent
	}
	if num.parent == nil {
		return
	}
	num = num.parent.parent
	if num.left.value != nil {
		*num.left.value += i
		return
	}
	num = num.left.ptr
	for num.right.ptr != nil {
		num = num.right.ptr
	}
	*num.right.value += i
}

func addToNextRight(num *snailfishNum, i int) {
	for isRightNode(num) {
		num = num.parent.parent
	}
	if num.parent == nil {
		return
	}
	num = num.parent.parent
	if num.right.value != nil {
		*num.right.value += i
		return
	}
	num = num.right.ptr
	for num.left.ptr != nil {
		num = num.left.ptr
	}
	*num.left.value += i
}

func isLeftNode(num *snailfishNum) bool {
	if num.parent != nil {
		return num.parent.parent.left.ptr == num
	}
	return false
}

func isRightNode(num *snailfishNum) bool {
	if num.parent != nil {
		return num.parent.parent.right.ptr == num
	}
	return false
}

func getSFNMagnitude(num *snailfishNum) int {
	return 3*getSFEMagnitude(num.left) + 2*getSFEMagnitude(num.right)
}

func getSFEMagnitude(elem *snailfishElem) int {
	if elem.value != nil {
		return *elem.value
	} else {
		return getSFNMagnitude(elem.ptr)
	}
}
