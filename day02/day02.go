package day02

import (
	"bufio"
	"fmt"
	"os"

	"github.com/winchest3r/aoc2024/utils"
)

func SafeReport(arr []int) bool {
	if len(arr) < 2 {
		return true
	}
	sign := utils.Sign(arr[0] - arr[1])
	for i := 1; i < len(arr); i++ {
		diff := arr[i-1] - arr[i]
		if sign != utils.Sign(diff) || diff == 0 || utils.AbsInt(diff) > 3 {
			return false
		}
	}
	return true
}

func SafeReportDampener(arr []int) bool {
	sign, prev, i := utils.Sign(arr[0]-arr[1]), arr[0], 1
	removed := false
	if utils.AbsInt(arr[0]-arr[1]) > 3 || arr[0]-arr[1] == 0 {
		removed = true
		i = 2
		sign = utils.Sign(arr[0] - arr[2])
	}
	for _, val := range arr[i:] {
		diff := prev - val
		if sign != utils.Sign(diff) || diff == 0 || utils.AbsInt(diff) > 3 {
			if removed {
				return false
			}
			removed = true
		} else {
			prev = val
		}
	}

	return true
}

func SafeReportDampenerSlow(arr []int) bool {
	if SafeReport(arr) {
		return true
	}

	// O(nl), n - rows, l - row size
	for i := 0; i < len(arr); i++ {
		newArr := make([]int, 0, len(arr)-1)
		newArr = append(newArr, arr[:i]...)
		newArr = append(newArr, arr[i+1:]...)
		if SafeReport(newArr) {
			return true
		}
	}

	return false
}

func ReadInputPartOne(fname string) [][]int {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, utils.FieldsToInt(scanner.Text()))
	}

	return result
}

func SolvePartOne(fname string) {
	data := ReadInputPartOne(fname)
	safe := 0
	for _, report := range data {
		if SafeReport(report) {
			safe++
		}
	}
	fmt.Println(safe)
}

func SolvePartTwo(fname string) {
	data := ReadInputPartOne(fname)
	safe := 0
	for _, report := range data {
		if SafeReportDampenerSlow(report) {
			safe++
		}
	}
	fmt.Println(safe)
}
