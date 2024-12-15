package day15

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestReadInput(t *testing.T) {
	_, got := ReadInput("data/example_small")
	want := []byte{'<', '^', '^', '>', '>', '>', 'v', 'v', '<', 'v', '>', '>', 'v', '<', '<'}
	assertion.AssertDeepEqual(t, got, want)
}

func TestCalculateGPS(t *testing.T) {
	mp, orders := ReadInput("data/example")
	MoveInOrder(mp, orders)
	got := CalculateGPS(mp)
	want := 10092
	assertion.AssertEqual(t, got, want)
}

func TestReadInputPartTwo(t *testing.T) {
	mp, _ := ReadInputPartTwo("data/example")
	mp.Print()
}
