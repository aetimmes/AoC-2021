package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	gamma := make([]int, len(scanner.Text()))
	aggregate(scanner.Text(), gamma)
	for scanner.Scan() {
		aggregate(scanner.Text(), gamma)
	}

	g, e := 0, 0

	for i, _ := range gamma {
		if gamma[len(gamma)-i-1] > 0 {
			g += int(math.Pow(2, float64(i)))
		} else {
			e += int(math.Pow(2, float64(i)))
		}
	}
	fmt.Println(g * e)
}

func aggregate(s string, gamma []int) {
	for i, c := range s {
		c, _ := strconv.Atoi(string(c))
		gamma[i] += c*2 - 1
	}
}
