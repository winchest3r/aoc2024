package day23

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestFindThreeCons(t *testing.T) {
	got := FindThreeCons(ReadInput("data/example"))
	want := 12
	assertion.AssertEqual(t, got, want)
}

func TestFindLargestNetwork(t *testing.T) {
	got := FindLargestNetwork(ReadInput("data/input"))
	want := []string{"co", "de", "ka", "ta"}
	assertion.AssertEqual(t, len(got), len(want))
}
