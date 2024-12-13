package day13

import (
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"

	"github.com/winchest3r/aoc2024/utils"
)

type Pair struct {
	X *big.Int
	Y *big.Int
}

type Machine struct {
	A     Pair
	B     Pair
	Prize Pair
}

func CalculateEquation(m Machine) (a, b *big.Int) {
	AyX := big.NewInt(0).Mul(m.A.Y, m.Prize.X)
	AxY := big.NewInt(0).Mul(m.A.X, m.Prize.Y)
	AyBx := big.NewInt(0).Mul(m.A.Y, m.B.X)
	AxBy := big.NewInt(0).Mul(m.A.X, m.B.Y)
	num := big.NewInt(0).Sub(AyX, AxY)
	denom := big.NewInt(0).Sub(AyBx, AxBy)
	b = big.NewInt(0).Div(num, denom) // amount of presses on button B
	Bxb := big.NewInt(0).Mul(m.B.X, b)
	num = big.NewInt(0).Sub(m.Prize.X, Bxb)
	a = big.NewInt(0).Div(num, m.A.X) // amount of presses on button A
	return
}

func IsRightSolution(m Machine, a, b *big.Int) bool {
	Axa := big.NewInt(0).Mul(m.A.X, a)
	Bxb := big.NewInt(0).Mul(m.B.X, b)
	Aya := big.NewInt(0).Mul(m.A.Y, a)
	Byb := big.NewInt(0).Mul(m.B.Y, b)
	x := big.NewInt(0).Add(Axa, Bxb)
	y := big.NewInt(0).Add(Aya, Byb)
	return x.Cmp(m.Prize.X) == 0 && y.Cmp(m.Prize.Y) == 0
}

func CalculateBestPrice(m Machine) *big.Int {
	a, b := CalculateEquation(m)
	// need to check if we have decimal amount of presses on buttons
	if !IsRightSolution(m, a, b) {
		return nil
	}
	a3 := big.NewInt(0).Mul(big.NewInt(3), a)
	return big.NewInt(0).Add(a3, b)
}

func ReadInputPartOne(fname string) []Machine {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(data), "\r\n\r\n")
	pattern, _ := regexp.Compile(`Button A: X\+(\d+), Y\+(\d+)\r\nButton B: X\+(\d+), Y\+(\d+)\r\nPrize: X=(\d+), Y=(\d+)`)
	result := make([]Machine, 0)
	for _, text := range arr {
		m := pattern.FindStringSubmatch(text)
		vals := utils.CastToBigInt(m[1:])
		result = append(result, Machine{
			A:     Pair{vals[0], vals[1]},
			B:     Pair{vals[2], vals[3]},
			Prize: Pair{vals[4], vals[5]},
		})
	}
	return result
}

func Add1e13(i *big.Int) *big.Int {
	value, _ := big.NewInt(0).SetString("10000000000000", 10)
	return big.NewInt(0).Add(value, i)
}

func ReadInputPartTwo(fname string) []Machine {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(data), "\r\n\r\n")
	pattern, _ := regexp.Compile(`Button A: X\+(\d+), Y\+(\d+)\r\nButton B: X\+(\d+), Y\+(\d+)\r\nPrize: X=(\d+), Y=(\d+)`)
	result := make([]Machine, 0)
	for _, text := range arr {
		m := pattern.FindStringSubmatch(text)
		vals := utils.CastToBigInt(m[1:])
		result = append(result, Machine{
			A:     Pair{vals[0], vals[1]},
			B:     Pair{vals[2], vals[3]},
			Prize: Pair{Add1e13(vals[4]), Add1e13(vals[5])},
		})
	}
	return result
}

func SolvePartOne(fname string) {
	machines := ReadInputPartOne(fname)
	result := big.NewInt(0)
	for _, m := range machines {
		price := CalculateBestPrice(m)
		if price == nil {
			continue
		}
		newRes := big.NewInt(0).Add(result, price)
		result = newRes
	}
	fmt.Println(result)
}

func SolvePartTwo(fname string) {
	machines := ReadInputPartTwo(fname)
	result := big.NewInt(0)
	for _, m := range machines {
		price := CalculateBestPrice(m)
		if price == nil {
			continue
		}
		newRes := big.NewInt(0).Add(result, price)
		result = newRes
	}
	fmt.Println(result)
}
