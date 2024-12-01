package utils

import "testing"

func AssertEqual[T comparable](tb testing.TB, got, want T) {
	tb.Helper()
	if got != want {
		tb.Errorf("got '%v' want '%v'", got, want)
	}
}

func AssertEqualFunc[T any](tb testing.TB, got, want T, pred func(a, b T) bool) {
	tb.Helper()
	if !pred(got, want) {
		tb.Errorf("got '%v' want '%v'", got, want)
	}
}
