package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x, y, aim := 0, 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		tokens := strings.Fields(current)
		magnitude, _ := strconv.Atoi(tokens[1])
		switch dir := tokens[0]; dir {
		case "forward":
			x += magnitude
			y += magnitude * aim
		case "up":
			aim -= magnitude
		case "down":
			aim += magnitude
		}

	}
	fmt.Println(x * y)
}
