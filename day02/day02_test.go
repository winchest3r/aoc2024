package day02

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestSafeReport(t *testing.T) {
	t.Run("7 6 4 2 1 safe", func(t *testing.T) {
		got := SafeReport([]int{7, 6, 4, 2, 1})
		want := true
		assertion.AssertEqual(t, got, want)
	})
	t.Run("1 2 7 8 9 unsafe", func(t *testing.T) {
		got := SafeReport([]int{1, 2, 7, 8, 9})
		want := false
		assertion.AssertEqual(t, got, want)
	})
	t.Run("9 7 6 2 1 unsafe", func(t *testing.T) {
		got := SafeReport([]int{9, 7, 6, 2, 1})
		want := false
		assertion.AssertEqual(t, got, want)
	})
	t.Run("1 3 2 4 5 unsafe", func(t *testing.T) {
		got := SafeReport([]int{1, 3, 2, 4, 5})
		want := false
		assertion.AssertEqual(t, got, want)
	})
	t.Run("8 6 4 4 1 unsafe", func(t *testing.T) {
		got := SafeReport([]int{8, 6, 4, 4, 1})
		want := false
		assertion.AssertEqual(t, got, want)
	})
	t.Run("1 3 6 7 9 safe", func(t *testing.T) {
		got := SafeReport([]int{1, 3, 6, 7, 9})
		want := true
		assertion.AssertEqual(t, got, want)
	})
}

func TestSafeReportDampener(t *testing.T) {
	t.Run("7 6 4 2 1 safe", func(t *testing.T) {
		got := SafeReportDampener([]int{7, 6, 4, 2, 1})
		want := true
		assertion.AssertEqual(t, got, want)
	})
	t.Run("1 2 7 8 9 unsafe", func(t *testing.T) {
		got := SafeReportDampener([]int{1, 2, 7, 8, 9})
		want := false
		assertion.AssertEqual(t, got, want)
	})
	t.Run("9 7 6 2 1 unsafe", func(t *testing.T) {
		got := SafeReportDampener([]int{9, 7, 6, 2, 1})
		want := false
		assertion.AssertEqual(t, got, want)
	})
	t.Run("1 3 2 4 5 safe", func(t *testing.T) {
		got := SafeReportDampener([]int{1, 3, 2, 4, 5})
		want := true
		assertion.AssertEqual(t, got, want)
	})
	t.Run("8 6 4 4 1 safe", func(t *testing.T) {
		got := SafeReportDampener([]int{8, 6, 4, 4, 1})
		want := true
		assertion.AssertEqual(t, got, want)
	})
	t.Run("1 3 6 7 9 safe", func(t *testing.T) {
		got := SafeReportDampener([]int{1, 3, 6, 7, 9})
		want := true
		assertion.AssertEqual(t, got, want)
	})
}
