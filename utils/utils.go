package utils

import (
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
