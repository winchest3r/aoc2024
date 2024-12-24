package day24

import (
	"fmt"
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestGetCombinedNumber(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		gates, ops := ReadInput("data/example")
		Process(gates, ops)
		got := GetCombinedNumber(gates, "z")
		want := 4
		assertion.AssertEqual(t, got, want)
	})
	t.Run("example2", func(t *testing.T) {
		gates, ops := ReadInput("data/example2")
		Process(gates, ops)
		got := GetCombinedNumber(gates, "z")
		want := 2024
		assertion.AssertEqual(t, got, want)
	})
}

func TestExample3(t *testing.T) {
	gates, ops := ReadInput("data/example3")
	ops[0].C, ops[5].C = ops[5].C, ops[0].C
	ops[1].C, ops[2].C = ops[2].C, ops[1].C
	Process(gates, ops)
	x := GetCombinedNumber(gates, "x")
	y := GetCombinedNumber(gates, "y")
	z := GetCombinedNumber(gates, "z")
	fmt.Println(x, y, z)
	assertion.AssertEqual(t, z, x+y)
}
