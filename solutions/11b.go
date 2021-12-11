package solutions

func F11b(input string) int {
	octopi, r, c := parseGrid(input)
	for i := 1; ; i++ {
		if processFlashes(&octopi, r, c) == r*c {
			return i
		}
	}
}
