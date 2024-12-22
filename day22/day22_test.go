package day22

import (
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestCalculateNextSecrets(t *testing.T) {
	got := CalculateNSecretsWithDiff(123, 10)
	want := []Pair{
		{15887950, -3},
		{16495136, 6},
		{527345, -1},
		{704524, -1},
		{1553684, 0},
		{12683156, 2},
		{11100544, -2},
		{12249484, 0},
		{7753432, -2},
		{5908254, 2},
	}
	assertion.AssertDeepEqual(t, got, want)
}

func TestCalculateNBestChanges(t *testing.T) {
	got := CalculateNBestChanges(123, 10, 1)
	want := []Price{
		{6, []int{-1, -1, 0, 2}},
	}
	assertion.AssertDeepEqual(t, got, want)
}

func TestCalculateNextSecret(t *testing.T) {
	got := CalculateNSecret(1, 2000)
	want := 8685429
	assertion.AssertEqual(t, got, want)
}

func TestReadInput(t *testing.T) {
	got := ReadInput("data/example")
	want := []int{
		1,
		10,
		100,
		2024,
	}
	assertion.AssertDeepEqual(t, got, want)
}

func TestGetBestPrice(t *testing.T) {
	buyers := ReadInput("data/example2")
	got := GetBestPrice(buyers, 2000)
	want := 23
	assertion.AssertEqual(t, got, want)
}
