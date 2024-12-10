package day10

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestFindScore(t *testing.T) {
	fname := "data/example_small"
	data := ReadInput(fname)
	got, _ := FindScore(data, Pair{0, 0}, 1, 9)
	want := 1
	assertion.AssertEqual(t, got, want)
}

func TestFindScorePartTwo(t *testing.T) {
	fname := "data/example_small"
	data := ReadInput(fname)
	_, got := FindScore(data, Pair{0, 0}, 1, 9)
	want := 16
	assertion.AssertEqual(t, got, want)
}
