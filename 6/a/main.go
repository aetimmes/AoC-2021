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
	fish := AStoi(strings.Split(scanner.Text(), ","))

	for i := 0; i < 80; i++ {
		for j := range fish {
			if fish[j] == 0 {
				fish[j] = 6
				fish = append(fish, 8)
			} else {
				fish[j] -= 1
			}
		}
	}
	fmt.Println(len(fish))
}

func AStoi(num_strings []string) []int {
	result := make([]int, 0, len(num_strings))
	for i := range num_strings {
		temp, _ := strconv.Atoi(num_strings[i])
		result = append(result, temp)
	}
	return result
}
