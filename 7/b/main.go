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

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	positions := AStof(strings.Split(scanner.Text(), ","))
	sort.Float64s(positions)

	midpoint := (averageFloat64(positions))
	result := math.Min(
		calcCrabFuel(positions, math.Floor(midpoint)),
		calcCrabFuel(positions, math.Ceil(midpoint)),
	)
	fmt.Println(int(result))
}

func AStof(num_strings []string) []float64 {
	result := make([]float64, 0, len(num_strings))
	for i := range num_strings {
		temp, _ := strconv.ParseFloat(num_strings[i], 64)
		result = append(result, temp)
	}
	return result
}

func averageFloat64(f []float64) float64 {
	result := 0.
	for i := range f {
		result += f[i]
	}
	return result / float64(len(f))
}

func calcCrabFuel(positions []float64, midpoint float64) float64 {
	result := 0.
	for i := range positions {
		diff := math.Abs(positions[i] - midpoint)
		result += (diff + 1) * diff / 2
	}
	return result
}
