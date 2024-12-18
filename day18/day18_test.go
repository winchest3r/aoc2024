package day18

import (
	"fmt"
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestBfs(t *testing.T) {
	data := ReadInput("data/example")
	got := Bfs(data, Pair{0, 0}, Pair{6, 6}, 12)
	want := 22
	assertion.AssertEqual(t, got, want)
}

func TestBfsPartTwo(t *testing.T) {
	data := ReadInput("data/example")
	got := ""
	for i := 12; i < len(data); i++ {
		val := Bfs(data, Pair{0, 0}, Pair{6, 6}, i)
		if val == 0 {
			got = fmt.Sprintf("%d,%d", data[i-1].X, data[i-1].Y)
			break
		}
	}
	want := "6,1"
	assertion.AssertEqual(t, got, want)
}
