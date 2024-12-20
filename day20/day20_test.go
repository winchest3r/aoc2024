package day20

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestGetBestFairPath(t *testing.T) {
	data := ReadInput("data/example")
	got := len(GetBestFairPath(data))
	want := 85
	assertion.AssertEqual(t, got, want)
}

func TestGetBestTime(t *testing.T) {
	data := ReadInput("data/example")
	got := GetBestTime(data, data.Start)
	want := 84
	assertion.AssertEqual(t, got, want)
}

func TestCountWithCheats(t *testing.T) {
	t.Run("two-seconds", func(t *testing.T) {
		data := ReadInput("data/example")
		path := GetBestFairPath(data)
		got := CountWithCheats(data, path, 1, 2)
		want := 44
		assertion.AssertEqual(t, got, want)
	})
	t.Run("twenty-seconds", func(t *testing.T) {
		data := ReadInput("data/example")
		path := GetBestFairPath(data)
		got := CountWithCheats(data, path, 50, 20)
		want := 285
		assertion.AssertEqual(t, got, want)
	})
}
