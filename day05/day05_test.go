package day05

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestGetRightUpdates(t *testing.T) {
	got := GetRightUpdates(ReadInput("data/example"))
	want := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
	}
	assertion.AssertDeepEqual(t, got, want)
}

func TestGetWrongUpdates(t *testing.T) {
	got := GetWrongUpdates(ReadInput("data/example"))
	want := [][]int{
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	assertion.AssertDeepEqual(t, got, want)
}

func TestGetReorderedUpdates(t *testing.T) {
	data := [][]int{
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	deps, _ := ReadInput("data/example")
	got := GetReorderedUpdates(deps, data)
	want := [][]int{
		{97, 75, 47, 61, 53},
		{61, 29, 13},
		{97, 75, 47, 29, 13},
	}
	assertion.AssertDeepEqual(t, got, want)
}

func TestGetMiddlePageNumber(t *testing.T) {
	data := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
	}
	got := GetMiddlePageNumber(data)
	want := 143
	assertion.AssertEqual(t, got, want)
}

func TestReadInput(t *testing.T) {
	_, updates := ReadInput("data/example")
	want := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	assertion.AssertDeepEqual(t, updates, want)
}
