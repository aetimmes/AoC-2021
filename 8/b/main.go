package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		l, r := parseLine(line)
		mappings := determineMappings(l)
		result += getValue(r, mappings)
	}

	fmt.Println(result)
}

func getValue(r []map[string]bool, mappings []map[string]bool) int {
	result := 0
	for i := range r {
		for j := range mappings {
			if reflect.DeepEqual(r[i], mappings[j]) {
				temp := j
				for k := i + 1; k < len(r); k++ {
					temp *= 10
				}
				result += temp
				break
			}
		}
	}
	return result
}

func parseLine(s string) ([]map[string]bool, []map[string]bool) {
	temp := strings.Split(s, "|")
	l := make([]map[string]bool, 0, 10)
	tl := strings.Fields(temp[0])
	for i := range tl {
		l = append(l, strToMap(tl[i]))
	}

	r := make([]map[string]bool, 0, 4)
	tr := strings.Fields(temp[1])
	for i := range tr {
		r = append(r, strToMap(tr[i]))
	}
	return l, r
}

func determineMappings(l []map[string]bool) []map[string]bool {
	result := make([]map[string]bool, 10)
	fives := make([]map[string]bool, 0, 3)
	sixes := make([]map[string]bool, 0, 3)
	for i := range l {
		c := l[i]
		switch len(c) {
		case 2:
			result[1] = c
		case 3:
			result[7] = c
		case 4:
			result[4] = c
		case 7:
			result[8] = c
		case 5:
			fives = append(fives, c)
		case 6:
			sixes = append(sixes, c)
		}
	}
	// Sixes: 6, 9, 0
	for i := range sixes {
		one, four, seven := 0, 0, 0
		for j := range sixes[i] {
			if result[1][j] == true {
				one += 1
			}
			if result[4][j] == true {
				four += 1
			}
			if result[7][j] == true {
				seven += 1
			}
		}
		if one == 1 {
			result[6] = sixes[i]
		} else if four == 4 {
			result[9] = sixes[i]
		} else {
			result[0] = sixes[i]
		}
	}
	// Fives: 2, 3, 5
	for i := range fives {
		one, four := 0, 0
		for j := range fives[i] {
			if result[1][j] == true {
				one += 1
			}
			if result[4][j] == true {
				four += 1
			}
		}
		if one == 2 {
			result[3] = fives[i]
		} else if four == 3 {
			result[5] = fives[i]
		} else {
			result[2] = fives[i]
		}
	}
	return result
}

func strToMap(s string) map[string]bool {
	result := make(map[string]bool)
	t := strings.Split(s, "")
	for i := range t {
		result[t[i]] = true
	}
	return result
}
