package solutions

import (
	"testing"
)

func TestExplode(t *testing.T) {
	table := map[string]string{
		"[[[[[9,8],1],2],3],4]":                 "[[[[0,9],2],3],4]",
		"[7,[6,[5,[4,[3,2]]]]]":                 "[7,[6,[5,[7,0]]]]",
		"[[6,[5,[4,[3,2]]]],1]":                 "[[6,[5,[7,0]]],3]",
		"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]": "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
		"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]":     "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
	}
	for k, v := range table {
		t.Logf("Parsing %s", k)
		sfn, rem := parseSnailfishNum(k, nil)
		if rem != "" {
			t.Fatalf("Parsing string failed; had %s left over", rem)
		}
		s := sfnToString(sfn)
		if s != k {
			t.Fatalf("Parsing string failed; expected %s, got %s", k, s)
		}
		x := explodeSFN(sfn, 1)
		if !x {
			t.Fatalf("Failed to explode %s", s)
		}
		s = sfnToString(sfn)
		if s != v {
			t.Fatalf("Failure exploding num: expected %s, got %s", v, s)
		}
		t.Logf("Exploding %s passed", k)
	}
}
