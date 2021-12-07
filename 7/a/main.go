package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
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
	scanner.Scan()
	positions := AStof(strings.Split(scanner.Text(), ","))
	sort.Float64s(positions)

	median := positions[len(positions)/2]

	for i := range positions {
		result += int(math.Abs(positions[i] - median))
	}
	fmt.Println(result)
}

func AStof(num_strings []string) []float64 {
	result := make([]float64, 0, len(num_strings))
	for i := range num_strings {
		temp, _ := strconv.ParseFloat(num_strings[i], 64)
		result = append(result, temp)
	}
	return result
}
