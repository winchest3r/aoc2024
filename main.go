package main

import (
	"github.com/winchest3r/aoc2024/day01"
	"github.com/winchest3r/aoc2024/day02"
	"github.com/winchest3r/aoc2024/day03"
	"github.com/winchest3r/aoc2024/day04"
)

func SolveDayOne() {
	fname := "day01/data/input"
	day01.SolvePartOne(fname)
	day01.SolvePartTwo(fname)
}

func SolveDayTwo() {
	fname := "day02/data/input"
	day02.SolvePartOne(fname)
	day02.SolvePartTwo(fname)
}

func SolveDayThree() {
	fname := "day03/data/input"
	day03.SolvePartOne(fname)
	day03.SolvePartTwo(fname)
}

func SolveDayFour() {
	fname := "day04/data/input"
	day04.SolvePartOne(fname)
	day04.SolvePartTwo(fname)
}

func main() {
	// SolveDayOne()
	// SolveDayTwo()
	// SolveDayThree()
	SolveDayFour()
}
