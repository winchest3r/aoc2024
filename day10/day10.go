package day10

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type Pair struct {
	row int
	col int
}

func ReadInput(fname string) map[Pair]int {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make(map[Pair]int)
	sc := bufio.NewScanner(file)
	i := 0
	for sc.Scan() {
		for j, r := range sc.Text() {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			data[Pair{i, j}] = num
		}
		i++
	}
	return data
}

func FindScore(data map[Pair]int, start Pair, step, height int) (int, int) {
	seen := make(map[Pair]bool)
	seen[start] = true
	queue := list.New()
	queue.PushBack(start)
	endPoints := make(map[Pair]bool)
	endPointsDistinct := make([]Pair, 0)
	dirs := []Pair{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	for queue.Len() > 0 {
		p := queue.Remove(queue.Front()).(Pair)
		if data[p] == height {
			endPoints[p] = true
			endPointsDistinct = append(endPointsDistinct, p)
		} else {
			for _, dir := range dirs {
				newP := Pair{p.row + dir.row, p.col + dir.col}
				if data[newP] != data[p]+step {
					continue
				}
				queue.PushBack(newP)
			}
		}
	}
	return len(endPoints), len(endPointsDistinct)
}

func SolvePartOne(fname string) {
	result := 0
	data := ReadInput(fname)
	for p := range data {
		if data[p] == 0 {
			val, _ := FindScore(data, p, 1, 9)
			result += val
		}
	}
	fmt.Println(result)
}

func SolvePartTwo(fname string) {
	result := 0
	data := ReadInput(fname)
	for p := range data {
		if data[p] == 0 {
			_, val := FindScore(data, p, 1, 9)
			result += val
		}
	}
	fmt.Println(result)
}
