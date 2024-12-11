package day11

import (
	"bufio"
	"fmt"
	"math/big"
	"os"

	"github.com/winchest3r/aoc2024/utils"
)

func ReadInput(fname string) []*big.Int {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	l := make([]*big.Int, 0)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		for _, num := range utils.FieldsToInt(sc.Text()) {
			l = append(l, big.NewInt(int64(num)))
		}
	}
	return l
}

func DivideNumber(num *big.Int) []*big.Int {
	s := num.String()
	first, second := big.NewInt(0), big.NewInt(0)
	first, _ = first.SetString(s[:len(s)/2], 10)
	second, _ = second.SetString(s[len(s)/2:], 10)
	return []*big.Int{first, second}
}

func CountDigits(num *big.Int) int {
	counter := 0
	i := big.NewInt(0)
	i = i.Set(num)
	for i.Cmp(big.NewInt(0)) > 0 {
		i = i.Div(i, big.NewInt(10))
		counter++
	}
	return counter
}

func BlinkOptimized(values []*big.Int, steps int) *big.Int {
	if steps <= 0 {
		return big.NewInt(int64(len(values)))
	}
	total := big.NewInt(0)
	for _, val := range values {
		newValues := make([]*big.Int, 0)
		switch {
		case val.Cmp(big.NewInt(0)) == 0:
			newValues = append(newValues, big.NewInt(1))
		case CountDigits(val)%2 != 0:
			newVal := big.NewInt(0)
			newVal.Mul(val, big.NewInt(2024))
			newValues = append(newValues, newVal)
		default:
			newValues = append(newValues, DivideNumber(val)...)
		}
		total = total.Add(total, BlinkOptimized(newValues, steps-1))
	}
	return total
}

func SolvePartOne(fname string) {
	l := ReadInput(fname)
	fmt.Println(BlinkOptimized(l, 25))
}

func SolvePartTwo(fname string) {
	l := ReadInput(fname)
	fmt.Println(BlinkOptimized(l, 40))
}
