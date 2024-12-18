package day18

import (
	"bufio"
	"fmt"
	"os"
	"slices"

	"github.com/winchest3r/aoc2024/utils"
)

type Pair struct {
	X int
	Y int
}

func ReadInput(fname string) []Pair {
	file, _ := os.Open(fname)
	defer file.Close()
	bytes := make([]Pair, 0)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		b := utils.SplitToInt(sc.Text(), ",")
		bytes = append(bytes, Pair{b[0], b[1]})
	}
	return bytes
}

func Bfs(data []Pair, start, end Pair, sz int) int {
	seen := make(map[Pair]bool)
	score := make(map[Pair]int)
	queue := make([]Pair, 0, end.X*end.Y)
	queue = append(queue, start)
	seen[start] = true
	dirs := []Pair{
		{0, -1}, {1, 0}, {0, 1}, {-1, 0},
	}
	for len(queue) > 0 {
		p := queue[:1][0]
		queue = queue[1:]
		if p == end {
			break
		}
		for _, d := range dirs {
			newP := Pair{p.X + d.X, p.Y + d.Y}
			if seen[newP] || slices.Contains(data[:sz], newP) {
				continue
			}
			if newP.X < start.X || newP.X > end.X || newP.Y < start.Y || newP.Y > end.Y {
				continue
			}
			queue = append(queue, newP)
			seen[newP] = true
			// value contains score in X and time in Y
			score[newP] = score[p] + 1
		}
	}
	return score[end]
}

func SolvePartOne(fname string) {
	data := ReadInput(fname)
	fmt.Println(Bfs(data, Pair{0, 0}, Pair{70, 70}, 1024))
}

func SolvePartTwo(fname string) {
	data := ReadInput(fname)
	for i := 1024; i < len(data); i++ {
		if Bfs(data, Pair{0, 0}, Pair{70, 70}, i) == 0 {
			fmt.Printf("%d,%d\n", data[i-1].X, data[i-1].Y)
			break
		}
	}
}
