package day19

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(fname string) ([]string, []string) {
	file, _ := os.ReadFile(fname)
	data := strings.Split(string(file), "\r\n\r\n")
	patterns := strings.Split(data[0], ", ")
	designs := strings.Fields(data[1])
	return patterns, designs
}

func PatternSet(data []string) map[string]bool {
	res := make(map[string]bool)
	for _, v := range data {
		res[v] = true
	}
	return res
}

func IsValidDesign(patterns map[string]bool, design string) bool {
	mem := make(map[string]bool)
	return Rec(patterns, mem, design)
}

func Rec(patterns, mem map[string]bool, design string) bool {
	if mem[design] {
		return mem[design]
	}
	for p := range patterns {
		if len(p) <= len(design) && design[:len(p)] == p {
			mem[design] = mem[design] || Rec(patterns, mem, design[len(p):])
		}
	}
	if len(design) == 0 {
		return true
	}
	return mem[design]
}

func FindValidDesings(pat, des []string) int {
	res := 0
	pSet := PatternSet(pat)
	for _, d := range des {
		if IsValidDesign(pSet, d) {
			res += 1
		}
	}
	return res
}

func SolvePartOne(fname string) {
	p, d := ReadInput(fname)
	fmt.Println(FindValidDesings(p, d))
}

func SolvePartTwo(fname string) {

}
