package day1

import (
	"slices"
	"testing"

	"github.com/winchest3r/aoc2024/utils"
)

func TestSumOfDistances(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := SumOfDistances([]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3})
		want := 11
		utils.AssertEqual(t, got, want)
	})
}

func TestSimilarityScore(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := SimilarityScore([]int{3, 4, 2, 1, 3, 3}, map[int]int{4: 1, 3: 3, 5: 1, 9: 1})
		want := 31
		utils.AssertEqual(t, got, want)
	})
}

func TestReadInputPartOne(t *testing.T) {
	gotLeft, gotRight := ReadInputPartOne("data/example")
	wantLeft, wantRight := []int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}

	utils.AssertEqualFunc(t, gotLeft, wantLeft, slices.Equal)
	utils.AssertEqualFunc(t, gotRight, wantRight, slices.Equal)
}
