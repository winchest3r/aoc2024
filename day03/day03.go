package day03

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/winchest3r/aoc2024/utils"
)

func ParseMultiplyString(text string) int {
	pattern, err := regexp.Compile(`mul\((\d+),(\d+)\)`)
	if err != nil {
		panic(err)
	}

	var result int
	for _, s := range pattern.FindAllString(text, -1) {
		result += ParseMultiplyOperator(pattern, s)
	}

	return result
}

func ParseMultiplyStringExtended(text string) int {
	p, err := regexp.Compile(`mul\((\d+),(\d+)\)|don't\(\)|do\(\)`)
	if err != nil {
		panic(err)
	}

	var result int
	activated := true
	for _, s := range p.FindAllString(text, -1) {
		switch s {
		case "do()":
			activated = true
		case "don't()":
			activated = false
		default:
			if activated {
				result += ParseMultiplyOperator(p, s)
			}
		}
	}

	return result
}

func ParseMultiplyOperator(p *regexp.Regexp, s string) int {
	matches := p.FindStringSubmatch(s)
	nums := utils.FieldsToInt(strings.Join(matches[1:3], " "))
	return nums[0] * nums[1]
}

func ReadInput(fname string) string {
	file, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	return string(file)
}

func SolvePartOne(fname string) {
	data := ReadInput(fname)
	fmt.Println(ParseMultiplyString(data))
}

func SolvePartTwo(fname string) {
	data := ReadInput(fname)
	fmt.Println(ParseMultiplyStringExtended(data))
}
