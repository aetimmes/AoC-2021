package main

import (
	"bufio"
	"fmt"
	"log"
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
	previous, err := strconv.Atoi(scanner.Text())
	result := 0
	for scanner.Scan() {
		temp := scanner.Text()
		current, _ := strconv.Atoi(temp)
		if current > previous {
			result += 1
		}
		previous = current
	}
	fmt.Println(result)
}
