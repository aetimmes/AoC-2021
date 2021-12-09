package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func F1a(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	previous, _ := strconv.Atoi(scanner.Text())
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
