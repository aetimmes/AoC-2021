package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func F10b(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0
	for scanner.Scan() {
	}
	fmt.Println(result)
}
