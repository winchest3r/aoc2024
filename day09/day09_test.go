package day09

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestCalculateChecksum(t *testing.T) {
	mem := ReadInput("data/example")
	FragmentDisk(mem)
	got := CalculateChecksum(mem)
	want := 1928
	assertion.AssertEqual(t, got, want)
}

func TestFragmentDiskTwo(t *testing.T) {
	mem := ReadInputPartTwo("data/example")
	FragmentDiskTwo(mem)
}

func TestCalculateChecksumTwo(t *testing.T) {
	mem := ReadInputPartTwo("data/example")
	FragmentDiskTwo(mem)
	got := CalculateChecksumTwo(mem)
	want := 2858
	assertion.AssertEqual(t, got, want)
}
