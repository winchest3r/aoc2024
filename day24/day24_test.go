package day24

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestGetCombinedNumber(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		gates, ops := ReadInput("data/example")
		Process(gates, ops)
		got := GetCombinedNumber(gates)
		want := 4
		assertion.AssertEqual(t, got, want)
	})
	t.Run("example2", func(t *testing.T) {
		gates, ops := ReadInput("data/example2")
		Process(gates, ops)
		got := GetCombinedNumber(gates)
		want := 2024
		assertion.AssertEqual(t, got, want)
	})
}
