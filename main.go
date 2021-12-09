package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"

	"github.com/aetimmes/AoC-2021/solutions"
)

var funcs = map[string]interface{}{
	"1a": solutions.F1a,
	"1b": solutions.F1b,
	"2a": solutions.F2a,
	"2b": solutions.F2b,
	"3a": solutions.F3a,
	"3b": solutions.F3b,
	"4a": solutions.F4a,
	"4b": solutions.F4b,
	"5a": solutions.F5a,
	"5b": solutions.F5b,
	"6a": solutions.F6a,
	"6b": solutions.F6b,
	"7a": solutions.F7a,
	"7b": solutions.F7b,
	"8a": solutions.F8a,
	"8b": solutions.F8b,
	"9a": solutions.F9a,
	"9b": solutions.F9b,
}

func main() {
	test := flag.Bool("test", false, "use test input rather than primary input")
	flag.BoolVar(test, "t", false, "")
	flag.Parse()
	if flag.NArg() != 1 {
		err := fmt.Errorf("expected 1 positional argument, got %d", flag.NArg())
		fmt.Println(err)
		os.Exit(1)
	}
	v, ok := funcs[flag.Arg(0)]
	if !ok {
		err := fmt.Errorf("no solution function for problem %s", flag.Arg(0))
		fmt.Println(err)
		os.Exit(1)
	}
	function := reflect.ValueOf(v)

	filename := string(flag.Arg(0)[0]) + ".txt"
	if *test {
		filename = "test-" + filename
	}
	path := "inputs/" + filename
	params := make([]reflect.Value, 0, 1)
	params = append(params, reflect.ValueOf(path))
	function.Call(params)
}
