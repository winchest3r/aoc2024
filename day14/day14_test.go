package day14

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestReadInput(t *testing.T) {
	fname := "data/example_small"
	p, v := ReadInput(fname)
	wantP := []Pair{
		{0, 4}, {6, 3}, {10, 3},
	}
	wantV := []Pair{
		{3, -3}, {-1, -3}, {-1, 2},
	}
	assertion.AssertDeepEqual(t, p, wantP)
	assertion.AssertDeepEqual(t, v, wantV)
}

func TestSafetyFactor(t *testing.T) {
	fname := "data/example"
	p, v := ReadInput(fname)
	s := CreateSpace(p, v, Pair{11, 7})
	ProcessSpace(s, 100)
	got := s.SafetyFactor()
	want := 12
	assertion.AssertEqual(t, got, want)
}
