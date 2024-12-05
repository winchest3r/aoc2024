package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/winchest3r/aoc2024/utils"
)

func GetRightUpdates(deps map[int][]int, upd [][]int) [][]int {
	result := make([][]int, 0)
loop:
	for _, u := range upd {
		for i := len(u) - 2; i >= 0; i-- {
			arr, ok := deps[u[i+1]]
			if !ok || !slices.Contains(arr, u[i]) {
				continue loop
			}
		}
		result = append(result, u)
	}
	return result
}

func GetWrongUpdates(deps map[int][]int, upd [][]int) [][]int {
	result := make([][]int, 0)
	for _, u := range upd {
		for i := len(u) - 2; i >= 0; i-- {
			arr, ok := deps[u[i+1]]
			if !ok || !slices.Contains(arr, u[i]) {
				result = append(result, u)
				break
			}
		}
	}
	return result
}

func DepPath(deps map[int][]int) map[int][]int {
	path := make(map[int][]int)
	for key := range deps {
		for _, val := range deps[key] {
			if !slices.Contains(path[val], key) {
				path[val] = append(path[val], key)
			}
		}
	}
	return path
}

func GetReorderedUpdates(deps map[int][]int, upd [][]int) [][]int {
	path := DepPath(deps)
	result := make([][]int, 0)
	for _, u := range upd {
		result = append(result, Reorder(path, u))
	}
	return result
}

func Reorder(path map[int][]int, upd []int) []int {
	result := make([]int, 0, len(upd))
	seen := make(map[int]bool)

	var dfs func(int)
	dfs = func(num int) {
		// filter seen and not existed numbers
		if seen[num] || !slices.Contains(upd, num) {
			return
		}
		seen[num] = true

		for _, p := range path[num] {
			dfs(p)
		}
		result = append(result, num)
	}

	for _, num := range upd {
		dfs(num)
	}

	slices.Reverse(result)

	return result
}

func GetMiddlePageNumber(upd [][]int) int {
	result := 0
	for _, row := range upd {
		result += row[len(row)/2]
	}
	return result
}

func ReadInput(fname string) (map[int][]int, [][]int) {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	deps := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		arr := utils.SplitToInt(line, "|")
		deps[arr[1]] = append(deps[arr[1]], arr[0])
	}
	updates := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		arr := utils.SplitToInt(line, ",")
		updates = append(updates, arr)
	}
	return deps, updates
}

func SolvePartOne(fname string) {
	rightUpdates := GetRightUpdates(ReadInput(fname))
	fmt.Println(GetMiddlePageNumber(rightUpdates))
}

func SolvePartTwo(fname string) {
	deps, upds := ReadInput(fname)
	wrongUpdates := GetWrongUpdates(deps, upds)
	reordered := GetReorderedUpdates(deps, wrongUpdates)
	fmt.Println(GetMiddlePageNumber(reordered))
}
