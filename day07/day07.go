package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/winchest3r/aoc2024/utils"
)

type Value struct {
	Result  int
	Numbers []int
}

func ReadInput(fname string) []Value {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	values := make([]Value, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": ")
		res, err := strconv.Atoi(line[0])
		if err != nil {
			panic(err)
		}
		nums := utils.FieldsToInt(line[1])
		values = append(values, Value{res, nums})
	}
	return values
}

func GenerateOps(v Value) []string {
	ops := make([]string, 0)
	var gen func(int, string)
	gen = func(i int, comb string) {
		if len(comb) == len(v.Numbers) {
			ops = append(ops, comb)
			return
		}
		gen(i+1, comb+"+")
		gen(i+1, comb+"*")
	}
	gen(0, "")
	return ops
}

func GenerateOpsPartTwo(v Value) []string {
	ops := make([]string, 0)
	var gen func(int, string)
	gen = func(i int, comb string) {
		if len(comb) == len(v.Numbers) {
			ops = append(ops, comb)
			return
		}
		gen(i+1, comb+"+")
		gen(i+1, comb+"*")
		gen(i+1, comb+"|")
	}
	gen(0, "")
	return ops
}

func Reduce(nums []int, ops string) int {
	result := nums[0]
	for i, n := range nums[1:] {
		switch ops[i] {
		case '+':
			result += n
		case '*':
			result *= n
		case '|':
			newNum, err := strconv.Atoi(fmt.Sprintf("%d%d", result, n))
			if err != nil {
				panic(err)
			}
			result = newNum
		}
	}
	return result
}

func RightEquation(v Value) bool {
	ops := GenerateOps(v)
	for _, op := range ops {
		if v.Result == Reduce(v.Numbers, op) {
			return true
		}
	}
	return false
}

func RightEquationPartTwo(v Value) bool {
	ops := GenerateOpsPartTwo(v)
	for _, op := range ops {
		if v.Result == Reduce(v.Numbers, op) {
			return true
		}
	}
	return false
}

func SolvePartOne(fname string) {
	vals := ReadInput(fname)
	result := 0
	for _, val := range vals {
		if RightEquation(val) {
			result += val.Result
		}
	}
	fmt.Println(result)
}

func SolvePartTwo(fname string) {
	vals := ReadInput(fname)
	result := 0
	for _, val := range vals {
		if RightEquationPartTwo(val) {
			result += val.Result
		}
	}
	fmt.Println(result)
}
