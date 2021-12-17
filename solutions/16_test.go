package solutions

import (
	"testing"
)

func TestParse(t *testing.T) {
	start := 1000000
	step := 100000
	for i := start; i < start+step; i++ {
		bits := makeBits(i)
		a, _, _ := parseLiteral(bits)
		if *a != uint64(i) {
			t.Errorf("%d failed", i)
		}
	}
}

func makeBits(i int) []bit {
	temp := make([]bit, 0)
	for i > 0 {
		temp = append(temp, i%2 == 1)
		i >>= 1
	}
	for len(temp)%4 != 0 {
		temp = append(temp, false)
	}
	result := make([]bit, 0)
	for i := len(temp) - 1; i >= 0; i-- {
		if i%4 == 3 {
			result = append(result, i != 3)
		}
		result = append(result, temp[i])
	}
	return result
}
