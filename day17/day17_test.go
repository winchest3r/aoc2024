package day17

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestReadInput(t *testing.T) {
	fname := "data/example"
	got := ReadInput(fname)
	want := NewMachine(729, 0, 0, []int{0, 1, 5, 4, 3, 0})
	assertion.AssertDeepEqual(t, got, want)
}

func TestProcessAll(t *testing.T) {
	fname := "data/example"
	m := ReadInput(fname)
	m.ProcessAll()
	got := m.OutputJoin()
	want := "4,6,3,5,6,3,5,2,1,0"
	assertion.AssertEqual(t, got, want)
}

func TestFindCopy(t *testing.T) {
	fname := "data/example2"
	m := ReadInput(fname)
	got := FindCopy(m)
	want := 117440
	assertion.AssertEqual(t, got, want)
}
