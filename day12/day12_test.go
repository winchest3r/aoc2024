package day12

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestDivideToRegions(t *testing.T) {
	t.Run("small", func(t *testing.T) {
		fname := "data/example_small"
		data := ReadInput(fname)
		got := len(DivideToRegions(data))
		want := 5
		assertion.AssertEqual(t, got, want)
	})
	t.Run("example", func(t *testing.T) {
		fname := "data/example"
		data := ReadInput(fname)
		got := len(DivideToRegions(data))
		want := 11
		assertion.AssertEqual(t, got, want)
	})
}
