package day03

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestParseMultiplyString(t *testing.T) {
	got := ParseMultiplyString("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	want := 161

	assertion.AssertEqual(t, got, want)
}
func TestParseMultiplyStringExtended(t *testing.T) {
	got := ParseMultiplyStringExtended("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))")
	want := 48

	assertion.AssertEqual(t, got, want)
}
