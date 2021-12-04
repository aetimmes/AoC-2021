package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	}
	fmt.Println(result)
}

type BingoBoard struct {
	nums map[int]bool
	r1   map[int]bool
	r2   map[int]bool
	r3   map[int]bool
	r4   map[int]bool
	r5   map[int]bool
	c1   map[int]bool
	c2   map[int]bool
	c3   map[int]bool
	c4   map[int]bool
	c5   map[int]bool
}

// This sucks

func NewBoard(nums []int) BingoBoard {

	board := {
		nums,
		nums[0:5],
		nums[5:10],
		nums[10:15],
		nums[15:20],
		nums[20:25],

	}
}
