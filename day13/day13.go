package day13

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/winchest3r/aoc2024/utils"
)

type Pair struct {
	X int
	Y int
}

type Machine struct {
	A     Pair
	B     Pair
	Prize Pair
}

func CalculateBestPrice(m Machine) int {
	minPrice := math.MaxInt
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			cur := Pair{i*m.A.X + j*m.B.X, i*m.A.Y + j*m.B.Y}
			if cur == m.Prize {
				minPrice = min(minPrice, i*3+j*1)
			}
		}
	}
	if minPrice == math.MaxInt {
		return 0
	}
	return minPrice
}

func ReadInput(fname string) []Machine {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	arr := strings.Split(string(data), "\r\n\r\n")
	pattern, _ := regexp.Compile(`Button A: X\+(\d+), Y\+(\d+)\r\nButton B: X\+(\d+), Y\+(\d+)\r\nPrize: X=(\d+), Y=(\d+)`)
	result := make([]Machine, 0)
	for _, text := range arr {
		m := pattern.FindStringSubmatch(text)
		vals := utils.CastToInt(m[1:])
		result = append(result, Machine{
			A:     Pair{vals[0], vals[1]},
			B:     Pair{vals[2], vals[3]},
			Prize: Pair{vals[4], vals[5]},
		})
	}
	return result
}

func SolvePartOne(fname string) {
	machines := ReadInput(fname)
	result := 0
	for _, m := range machines {
		price := CalculateBestPrice(m)
		result += price
	}
	fmt.Println(result)
}

func SolvePartTwo(fname string) {

}
