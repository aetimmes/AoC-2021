package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var scores = map[string]int{
	")": 3,
	"]": 57,
	"}": 1197,
	">": 25137,
}

var bracePairs = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
	">": "<",
}

func F10a(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			result += scoreLine(line)
		}
	}
	fmt.Println(result)
}

func scoreLine(s string) int {

}
