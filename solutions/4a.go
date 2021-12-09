package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func F4a(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	num_strings := strings.Split(scanner.Text(), ",")
	called_nums := AStoi(num_strings)

	boards := make([]BingoBoard, 0, 100)
	current := make([]int, 0, 25)
	for scanner.Scan() {
		if scanner.Text() != "" {
			current = append(current, AStoi(strings.Fields(scanner.Text()))...)
			if len(current) == 25 {
				boards = append(boards, NewBoard(current))
				current = make([]int, 0, 25)
			}
		}
	}

	for i := range called_nums {
		for j := range boards {
			boards[j] = callNumber(boards[j], called_nums[i])
			if checkWin(boards[j]) {
				fmt.Println(getScore(boards[j], called_nums[i]))
				return
			}
		}
	}
}
