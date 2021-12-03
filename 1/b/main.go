package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
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
	r := ring.New(3)

	for i := 0; i < 3; i++ {
		scanner.Scan()
		r.Value, _ = strconv.Atoi(scanner.Text())
		r = r.Next()
	}

	result := 0
	for scanner.Scan() {
		current, _ := strconv.Atoi(scanner.Text())
		if current > int(r.Value.(int)) {
			result += 1
		}
		r.Value = current
		r = r.Next()
	}
	fmt.Println(result)
}
