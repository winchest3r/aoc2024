package day04

import (
	"reflect"
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestXmasCounter(t *testing.T) {
	got := XmasCounter([][]byte{
		[]byte("MMMSXXMASM"),
		[]byte("MSAMXMSMSA"),
		[]byte("AMXSXMAAMM"),
		[]byte("MSAMASMSMX"),
		[]byte("XMASAMXAMM"),
		[]byte("XXAMMXXAMA"),
		[]byte("SMSMSASXSS"),
		[]byte("SAXAMASAAA"),
		[]byte("MAMMMXMMMM"),
		[]byte("MXMXAXMASX"),
	})
	want := 18

	assertion.AssertEqual(t, got, want)
}

func TestMasCounter(t *testing.T) {
	got := MasCounter([][]byte{
		[]byte("MMMSXXMASM"),
		[]byte("MSAMXMSMSA"),
		[]byte("AMXSXMAAMM"),
		[]byte("MSAMASMSMX"),
		[]byte("XMASAMXAMM"),
		[]byte("XXAMMXXAMA"),
		[]byte("SMSMSASXSS"),
		[]byte("SAXAMASAAA"),
		[]byte("MAMMMXMMMM"),
		[]byte("MXMXAXMASX"),
	})
	want := 9

	assertion.AssertEqual(t, got, want)
}

func TestReadInput(t *testing.T) {
	got := ReadInput("data/example")

	want := [][]byte{
		[]byte("MMMSXXMASM"),
		[]byte("MSAMXMSMSA"),
		[]byte("AMXSXMAAMM"),
		[]byte("MSAMASMSMX"),
		[]byte("XMASAMXAMM"),
		[]byte("XXAMMXXAMA"),
		[]byte("SMSMSASXSS"),
		[]byte("SAXAMASAAA"),
		[]byte("MAMMMXMMMM"),
		[]byte("MXMXAXMASX"),
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("wrong expected input")
	}
}
