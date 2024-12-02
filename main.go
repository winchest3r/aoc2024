package main

import (
	day1 "github.com/winchest3r/aoc2024/day01"
	"github.com/winchest3r/aoc2024/day02"
)

func SolveDayOne() {
	input := "day01/data/input"
	day1.SolvePartOne(input)
	day1.SolvePartTwo(input)
}

func SolveDayTwo() {
	input := "day02/data/input"
	day02.SolvePartOne(input)
	day02.SolvePartTwo(input)
}

func main() {
	// SolveDayOne()
	SolveDayTwo()
}
