package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func F2a(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	x, y := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current := scanner.Text()
		tokens := strings.Fields(current)
		magnitude, _ := strconv.Atoi(tokens[1])
		switch dir := tokens[0]; dir {
		case "forward":
			x += magnitude
		case "up":
			y -= magnitude
		case "down":
			y += magnitude
		}

	}
	fmt.Println(x * y)
}
