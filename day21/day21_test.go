package day21

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestMove(t *testing.T) {
	p := NewNumberKeypad()
	r := NewRobot(p, "A")
	got := r.Move("4")
	want := "^^<<"
	assertion.AssertDeepEqual(t, got, want)
}

func TestMakeSetOfMovements(t *testing.T) {
	p := NewNumberKeypad()
	r := NewRobot(p, "A")
	got := len(MakeSetOfMovements(r, "029A"))
	want := len("<A^A>^^AvvvA")
	assertion.AssertEqual(t, got, want)
}

func TestProcessRobots(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		t1, t2, t3 := NewNumberKeypad(), NewDirectKeypad(), NewDirectKeypad()
		r1, r2, r3 := NewRobot(t1, "A"), NewRobot(t2, "A"), NewRobot(t3, "A")
		got := ProcessRobots([]Robot{r1, r2, r3}, "029A")
		want := "<vA<AA>>^AvAA<^A>A<v<A>>^AvA^A<vA>^A<v<A>^A>AAvA^A<v<A>A>^AAAvA<^A>A"
		assertion.AssertEqual(t, len(got), len(want))
	})
	t.Run("second", func(t *testing.T) {
		got := ThreeRobots("980A")
		want := "<v<A>>^AAAvA^A<vA<AA>>^AvAA<^A>A<v<A>A>^AAAvA<^A>A<vA>^A<A>A"
		assertion.AssertEqual(t, len(got), len(want))
	})
	t.Run("third", func(t *testing.T) {
		got := ThreeRobots("179A")
		want := "<v<A>>^A<vA<A>>^AAvAA<^A>A<v<A>>^AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A"
		assertion.AssertEqual(t, len(got), len(want))
	})
	t.Run("two robots first", func(t *testing.T) {
		t1, t2 := NewNumberKeypad(), NewDirectKeypad()
		r1, r2 := NewRobot(t1, "A"), NewRobot(t2, "A")
		got := ProcessRobots([]Robot{r1, r2}, "029A")
		want := "v<<A>>^A<A>AvA<^AA>A<vAAA>^A"
		assertion.AssertEqual(t, len(got), len(want))
	})
}

func TestCalculateSum(t *testing.T) {
	codes := ReadInput("data/example")
	got := CalculateSum(codes)
	want := 126384
	assertion.AssertEqual(t, got, want)
}
