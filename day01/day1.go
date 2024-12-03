package day01

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func SumOfDistances(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := range left {
		sum += int(math.Abs(float64(left[i] - right[i])))
	}

	return sum
}

func SimilarityScore(left []int, right map[int]int) int {
	sum := 0
	for _, val := range left {
		sum += val * right[val]
	}
	return sum
}

func ReadInputPartOne(fname string) ([]int, []int) {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	left, right := []int{}, []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())
		leftVal, _ := strconv.Atoi(arr[0])
		rightVal, _ := strconv.Atoi(arr[1])
		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	return left, right
}

func ReadInputPartTwo(fname string) ([]int, map[int]int) {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	left, right := []int{}, make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		arr := strings.Fields(scanner.Text())

		leftVal, _ := strconv.Atoi(arr[0])
		rightVal, _ := strconv.Atoi(arr[1])

		left = append(left, leftVal)
		right[rightVal]++
	}

	return left, right
}

func SolvePartOne(fname string) {
	l, r := ReadInputPartOne(fname)
	result := SumOfDistances(l, r)
	fmt.Println(result)
}

func SolvePartTwo(fname string) {
	l, r := ReadInputPartTwo(fname)
	result := SimilarityScore(l, r)
	fmt.Println(result)
}
