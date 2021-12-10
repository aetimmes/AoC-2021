package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/aetimmes/AoC-2021/solutions"
	"github.com/aetimmes/go-aoc-client/aocclient"
)

var funcs = map[string]interface{}{
	"1a":  solutions.F1a,
	"1b":  solutions.F1b,
	"2a":  solutions.F2a,
	"2b":  solutions.F2b,
	"3a":  solutions.F3a,
	"3b":  solutions.F3b,
	"4a":  solutions.F4a,
	"4b":  solutions.F4b,
	"5a":  solutions.F5a,
	"5b":  solutions.F5b,
	"6a":  solutions.F6a,
	"6b":  solutions.F6b,
	"7a":  solutions.F7a,
	"7b":  solutions.F7b,
	"8a":  solutions.F8a,
	"8b":  solutions.F8b,
	"9a":  solutions.F9a,
	"9b":  solutions.F9b,
	"10a": solutions.F10a,
	"10b": solutions.F10b,
}

var levelMap = map[byte]int{
	'a': 1,
	'b': 2,
}

func main() {
	test := flag.Bool("test", false, "use test input rather than primary input")
	flag.BoolVar(test, "t", false, "")
	noSubmit := flag.Bool("dry-run", false, "doesn't submit answer to AoC")
	flag.BoolVar(test, "d", false, "")
	flag.Parse()
	if flag.NArg() != 1 {
		err := fmt.Errorf("expected 1 positional argument, got %d", flag.NArg())
		fmt.Println(err)
		os.Exit(1)
	}
	v, ok := funcs[flag.Arg(0)]
	if !ok {
		log.Fatalf("no solution function for problem %s", flag.Arg(0))
	}
	year := 2021
	day, err := strconv.Atoi(flag.Arg(0)[:len(flag.Arg(0))-1])
	if err != nil {
		log.Fatalf("Failed to parse day %s: %s", flag.Arg(0), err)
	}
	sf, err := os.ReadFile("sessionID.txt")
	if err != nil {
		log.Fatalf("Failed to get sessionID: %s", err)
	}
	sessionID := strings.TrimSpace(string(sf))
	level := levelMap[flag.Arg(0)[len(flag.Arg(0))-1]]
	function := reflect.ValueOf(v)
	var input string
	if *test {
		inputFile, err := os.ReadFile(fmt.Sprintf("inputs/test-%d.txt", day))
		if err != nil {
			log.Fatalf("Failed to read test input file: %s", err)
		}
		input = string(inputFile)
	} else {
		input, err = aocclient.GetInput(year, day, sessionID)
		if err != nil {
			log.Fatalf("Failed to get problem input: %s", err)
		}
	}

	params := make([]reflect.Value, 0, 1)
	params = append(params, reflect.ValueOf(input))
	results := function.Call(params)
	if len(results) > 1 {
		log.Fatalf("wrong number of values returned by function")
	}

	answer := int(results[0].Int())

	if *test {
		fmt.Println(answer)
	} else if *noSubmit {
		fmt.Println("Answer: ", answer, "Result not submitted")
	} else {
		response_type, err := aocclient.SubmitAnswer(year, day, level, answer, sessionID)
		if err != nil {
			log.Printf("Error submitting answer: %s\n", err)
		}
		fmt.Println("Answer: ", answer, "Result: ", aocclient.ResponseTypeMap[response_type])
	}
}
