package solutions

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func F7b(filename string) {
	file, err := os.Open(filename)
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
