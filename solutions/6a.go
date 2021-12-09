package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func F6a(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fish := AStoi(strings.Split(scanner.Text(), ","))

	for i := 0; i < 80; i++ {
		for j := range fish {
			if fish[j] == 0 {
				fish[j] = 6
				fish = append(fish, 8)
			} else {
				fish[j] -= 1
			}
		}
	}
	fmt.Println(len(fish))
}
