package day13

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestReadInput(t *testing.T) {
	fname := "data/example"
	ReadInput(fname)
}

func TestCalculateBestPrice(t *testing.T) {
	m := Machine{
		A:     Pair{94, 34},
		B:     Pair{22, 67},
		Prize: Pair{8400, 5400},
	}
	got := CalculateBestPrice(m)
	want := 280
	assertion.AssertEqual(t, got, want)
}
