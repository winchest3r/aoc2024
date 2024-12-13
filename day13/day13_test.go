package day13

import (
	"math/big"
	"testing"

	"github.com/winchest3r/aoc2024/utils/assertion"
)

func TestReadInput(t *testing.T) {
	fname := "data/example"
	ReadInputPartTwo(fname)
}

func TestCalculateBestPrice(t *testing.T) {
	m := Machine{
		A:     Pair{big.NewInt(94), big.NewInt(34)},
		B:     Pair{big.NewInt(22), big.NewInt(67)},
		Prize: Pair{big.NewInt(8400), big.NewInt(5400)},
	}
	got := CalculateBestPrice(m).String()
	want := "280"
	assertion.AssertEqual(t, got, want)
}

func TestAdd1e13(t *testing.T) {
	i := big.NewInt(42)
	got := Add1e13(i)
	want, _ := big.NewInt(0).SetString("10000000000042", 10)
	if i.Cmp(big.NewInt(42)) != 0 {
		t.Error("initial value is corrupted")
	}
	if got.Cmp(want) != 0 {
		t.Error("result is wrong")
	}
}
