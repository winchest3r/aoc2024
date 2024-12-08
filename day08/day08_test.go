package day08

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestReadInput(t *testing.T) {
	fname := "data/example_small"
	got := ReadInput(fname)
	want := [][]byte{
		{'.', '1', '.', '.'},
		{'.', '1', '1', '.'},
		{'.', '.', '.', '.'},
		{'.', '2', '.', '.'},
	}
	assertion.AssertDeepEqual(t, got, want)
}

func TestCalculateAntinodes(t *testing.T) {
	t.Run("small example", func(t *testing.T) {
		got := len(CalculateAntinodes(ReadInput("data/example_small")))
		want := 4
		assertion.AssertEqual(t, got, want)
	})
	t.Run("small example", func(t *testing.T) {
		got := len(CalculateAntinodes(ReadInput("data/example")))
		want := 14
		assertion.AssertEqual(t, got, want)
	})
}

func TestCalculateResonantAntinodes(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := len(CalculateResonantAntinodes(ReadInput("data/example")))
		want := 34
		assertion.AssertEqual(t, got, want)
	})
	t.Run("example2", func(t *testing.T) {
		got := len(CalculateResonantAntinodes(ReadInput("data/example2")))
		want := 9
		assertion.AssertEqual(t, got, want)
	})
}
