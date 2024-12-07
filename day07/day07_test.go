package day07

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestReadInput(t *testing.T) {
	got := ReadInput("data/example")
	want := []Value{
		{190, []int{10, 19}},
		{3267, []int{81, 40, 27}},
		{83, []int{17, 5}},
		{156, []int{15, 6}},
		{7290, []int{6, 8, 6, 15}},
		{161011, []int{16, 10, 13}},
		{192, []int{17, 8, 14}},
		{21037, []int{9, 7, 18, 13}},
		{292, []int{11, 6, 16, 20}},
	}
	assertion.AssertDeepEqual(t, got, want)
}

func TestGenerateOps(t *testing.T) {
	got := GenerateOps(Value{42, []int{1, 2}})
	want := []string{
		"++",
		"+*",
		"*+",
		"**",
	}
	assertion.AssertDeepEqual(t, got, want)
}

func TestRightEquation(t *testing.T) {
	t.Run("right", func(t *testing.T) {
		got := RightEquation(Value{3267, []int{81, 40, 27}})
		want := true
		assertion.AssertEqual(t, got, want)
	})
	t.Run("wrong", func(t *testing.T) {
		got := RightEquation(Value{21037, []int{9, 7, 18, 13}})
		want := false
		assertion.AssertEqual(t, got, want)
	})
}
