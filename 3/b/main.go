package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nums := []string{}
	for scanner.Scan() {
		if len(scanner.Text()) > 0 {
			nums = append(nums, scanner.Text())
		}
	}

	co2 := BAtoi(partition(nums, true, 0))
	o2 := BAtoi(partition(nums, false, 0))

	fmt.Println(co2 * o2)
}

func partition(nums []string, common bool, index int) string {
	if len(nums) == 1 {
		return nums[0]
	}

	ones := []string{}
	zeroes := []string{}

	for i := range nums {
		if string(nums[i][index]) == "0" {
			zeroes = append(zeroes, nums[i])
		} else {
			ones = append(ones, nums[i])
		}
	}
	if len(ones) < len(zeroes) != common {
		return partition(ones, common, index+1)
	} else {
		return partition(zeroes, common, index+1)
	}
}

func BAtoi(num string) int {
	result := 0
	for i := range num {
		if string(num[len(num)-i-1]) == "1" {
			result += int(math.Pow(2, float64(i)))
		}
	}
	return result
}
