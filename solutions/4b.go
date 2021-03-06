package solutions

import (
	"bufio"
	"strconv"
	"strings"
)

func F4b(input string) int {
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

	live_boards := map[int]bool{}
	for i := range boards {
		live_boards[i] = true
	}

	for i := range called_nums {
		for j := range live_boards {
			boards[j] = callNumber(boards[j], called_nums[i])
			if checkWin(boards[j]) {
				if len(live_boards) == 1 {
					return getScore(boards[j], called_nums[i])
				}
				delete(live_boards, j)
			}
		}
	}
	panic("No solution found")
}

type BingoBoard struct {
	nums    map[int]bool
	rows    []map[int]bool
	columns []map[int]bool
}

func NewBoard(numbers []int) BingoBoard {
	SIZE := 5
	rows := make([]map[int]bool, 0, SIZE)
	columns := make([]map[int]bool, 0, SIZE)
	for i := 0; i < SIZE; i++ {
		rows = append(rows, make(map[int]bool))
		columns = append(columns, make(map[int]bool))
	}
	nums := make(map[int]bool)
	for i := range numbers {
		nums[numbers[i]] = true
		rows[i/SIZE][numbers[i]] = true
		columns[i%SIZE][numbers[i]] = true
	}
	return BingoBoard{nums, rows, columns}
}

func callNumber(board BingoBoard, number int) BingoBoard {
	delete(board.nums, number)
	for i := range board.rows {
		delete(board.rows[i], number)
	}
	for i := range board.columns {
		delete(board.columns[i], number)
	}
	return board
}

func checkWin(board BingoBoard) bool {
	for i := range board.rows {
		if len(board.rows[i]) == 0 {
			return true
		}
	}
	for i := range board.columns {
		if len(board.columns[i]) == 0 {
			return true
		}
	}
	return false
}

func getScore(board BingoBoard, number int) int {
	sum := 0
	for i := range board.nums {
		sum += i
	}
	return sum * number
}

func AStoi(num_strings []string) []int {
	result := make([]int, 0, len(num_strings))
	for i := range num_strings {
		temp, _ := strconv.Atoi(num_strings[i])
		result = append(result, temp)
	}
	return result
}
