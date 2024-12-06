package day06

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestReadInput(t *testing.T) {
	fname := "data/small_example"
	gotMap, gotStartCell, gotStartDir := ReadInput(fname)
	wantMap := map[Coord]CellType{
		{0, 0}: Obstacle,
		{1, 0}: Empty,
		{2, 0}: Empty,
		{0, 1}: Empty,
		{1, 1}: Empty,
		{2, 1}: Obstacle,
		{0, 2}: Visited,
		{1, 2}: Empty,
		{2, 2}: Empty,
	}
	wantStartCell, wantStartDir := Coord{0, 2}, Dir{0, -1}
	assertion.AssertDeepEqual(t, gotMap, wantMap)
	assertion.AssertEqual(t, gotStartCell, wantStartCell)
	assertion.AssertEqual(t, gotStartDir, wantStartDir)
}

func TestGetVisitedCells(t *testing.T) {
	t.Run("small example", func(t *testing.T) {
		got := GetVisitedCells(ReadInput("data/small_example"))
		want := 4
		assertion.AssertEqual(t, got, want)
	})
	t.Run("example", func(t *testing.T) {
		got := GetVisitedCells(ReadInput("data/example"))
		want := 41
		assertion.AssertEqual(t, got, want)
	})
}

func TestGetLoopsWithObstruction(t *testing.T) {
	got := GetLoopsWithObstruction(ReadInput("data/example"))
	want := 6
	assertion.AssertEqual(t, got, want)
}
