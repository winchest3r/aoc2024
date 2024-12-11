package day11

import (
	"math/big"
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestBlinkTwentyFive(t *testing.T) {
	fname := "data/example"
	l := ReadInput(fname)
	got := BlinkOptimized(l, 25)
	want := big.NewInt(55312)
	assertion.AssertEqualFunc(t, got, want, func(a, b *big.Int) bool {
		return a.Cmp(b) == 0
	})
}
