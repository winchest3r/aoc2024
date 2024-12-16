package day16

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestCalculateBestPath(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		fname := "data/example"
		got := CalculateBestPath(ReadInput(fname))
		want := 7036
		assertion.AssertEqual(t, got, want)
	})
	t.Run("example2", func(t *testing.T) {
		fname := "data/example2"
		got := CalculateBestPath(ReadInput(fname))
		want := 11048
		assertion.AssertEqual(t, got, want)
	})
}
