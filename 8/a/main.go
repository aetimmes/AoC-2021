package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		_, r := parseLine(scanner.Text())
		for i := range r {
			switch len(r[i]) {
			case 2, 3, 4, 7:
				result += 1
			default:
			}
		}
	}
	fmt.Println(result)
}

func parseLine(s string) ([]string, []string) {
	temp := strings.Split(s, "|")
	l := strings.Fields(temp[0])
	r := strings.Fields(temp[1])
	return l, r
}
