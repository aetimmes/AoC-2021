package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	seen := make(map[point]int)
	for scanner.Scan() {
		start, finish := parseLine(scanner.Text())
		if start.x == finish.x {
			i, j := enforceAscending(start.y, finish.y)
			for p := newPoint(start.x, i); p.y <= j; p.y++ {
				seen[p] += 1
			}
		} else if start.y == finish.y {
			i, j := enforceAscending(start.x, finish.x)
			for p := newPoint(i, start.y); p.x <= j; p.x++ {
				seen[p] += 1
			}
		} else if math.Abs(float64(start.x)-float64(finish.x)) == math.Abs(float64(start.y)-float64(finish.y)) {
			x_dir := 1
			if start.x > finish.x {
				x_dir = -1
			}
			y_dir := 1
			if start.y > finish.y {
				y_dir = -1
			}
			for p := newPoint(start.x, start.y); p.x != finish.x; p.x, p.y = p.x+x_dir, p.y+y_dir {
				seen[p] += 1
			}
		} else {
			panic("Weird line")
		}
	}
	for k, v := range seen {
		fmt.Println(strconv.Itoa(k.x), strconv.Itoa(k.y), v)
		if v > 1 {
			result += 1
		}
	}
	fmt.Println(result)
}

func parseLine(s string) (point, point) {
	fmt.Println("Parsing: ", s)
	tokens := strings.Split(s, " -> ")
	fmt.Println("Tokens: ", tokens)
	return newPointFromString(tokens[0]), newPointFromString(tokens[1])
}

func newPointFromString(s string) point {
	tokens := strings.Split(s, ",")
	x, _ := strconv.Atoi(tokens[0])
	y, _ := strconv.Atoi(tokens[1])
	return newPoint(x, y)
}

func newPoint(x int, y int) point {
	fmt.Println("Creating new point with x: ", x, ", y: ", y)
	return point{x, y}
}

func enforceAscending(x int, y int) (int, int) {
	if x <= y {
		return x, y
	}
	return y, x
}
