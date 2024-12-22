package utils

import (
	"math/big"
	"strconv"
	"strings"
)

// AbsInt absolute value but for integers.
func AbsInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// Sign return 1 or -1 depends on the parameter sign. Zero returns 0.
func Sign(a int) int {
	if a == 0 {
		return 0
	}
	return a / AbsInt(a)
}

// If input contains a row with numbers divided by space(s),
// use it to create array of integers.
//
// "4 8 15 16 23 42" -> []int{4, 8, 15, 16, 23, 42}
//
// Panic if something goes wrong.
func FieldsToInt(row string) []int {
	fields := strings.Fields(row)
	result := make([]int, 0, len(fields))
	for _, field := range fields {
		val, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		result = append(result, val)
	}
	return result
}

func SplitToInt(row, sep string) []int {
	fields := strings.Split(row, sep)
	result := make([]int, 0, len(fields))
	for _, field := range fields {
		val, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		result = append(result, val)
	}
	return result
}

func CastToInt(arr []string) []int {
	result := make([]int, 0, len(arr))
	for _, s := range arr {
		val, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		result = append(result, val)
	}
	return result
}

func CastToInt64(arr []string) []int64 {
	result := make([]int64, 0, len(arr))
	for _, s := range arr {
		val, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		result = append(result, int64(val))
	}
	return result
}

func CastToBigInt(arr []string) []*big.Int {
	result := make([]*big.Int, 0, len(arr))
	for _, s := range arr {
		val, ok := big.NewInt(0).SetString(s, 10)
		if !ok {
			panic("can't process string: " + s)
		}
		result = append(result, val)
	}
	return result
}

func CastToString(arr []int) []string {
	result := make([]string, len(arr))
	for i, val := range arr {
		result[i] = strconv.Itoa(val)
	}
	return result
}
