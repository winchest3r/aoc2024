package day19

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestReadInput(t *testing.T) {
	p, d := ReadInput("data/example")
	wantP := []string{"r", "wr", "b", "g", "bwu", "rb", "gb", "br"}
	wantD := []string{
		"brwrr",
		"bggr",
		"gbbr",
		"rrbgbr",
		"ubwu",
		"bwurrg",
		"brgr",
		"bbrgwb",
	}
	assertion.AssertDeepEqual(t, p, wantP)
	assertion.AssertDeepEqual(t, d, wantD)
}

func TestExample(t *testing.T) {
	p, d := ReadInput("data/example")
	got := FindValidDesings(p, d)
	want := 6
	assertion.AssertEqual(t, got, want)
}
