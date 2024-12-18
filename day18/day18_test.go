package day18

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestBfs(t *testing.T) {
	data := ReadInput("data/example")
	got := Bfs(data, Pair{0, 0}, Pair{6, 6}, 12)
	want := 22
	assertion.AssertEqual(t, got, want)
}
