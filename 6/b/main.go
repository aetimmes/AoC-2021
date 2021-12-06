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

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	initial_fish := AStoi(strings.Split(scanner.Text(), ","))

	fish := make([]int, 9)
	for i := range(initial_fish) {
		fish[initial_fish[i]] += 1
	}

	for i := 0; i < 256; i++ {
		fish = []int{
			fish[1],
			fish[2],
			fish[3],
			fish[4],
			fish[5],
			fish[6],
			fish[7] + fish[0],
			fish[8],
			fish[0],
		}
	}

	result := 0
	for i := range(fish) {
		result += fish[i]
	}
	fmt.Println(result)
}

func AStoi(num_strings []string) []int {
	result := make([]int, 0, len(num_strings))
	for i := range num_strings {
		temp, _ := strconv.Atoi(num_strings[i])
		result = append(result, temp)
	}
	return result
}
