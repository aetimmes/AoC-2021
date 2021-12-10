package solutions

import (
	"bufio"
	"strings"
)

func F4a(input string) int {
	scanner := bufio.NewScanner(strings.NewReader(input))
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
				return getScore(boards[j], called_nums[i])
			}
		}
	}
	panic("No solution found")
}
