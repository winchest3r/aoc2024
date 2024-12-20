package day20

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"slices"

	"github.com/winchest3r/aoc2024/utils"
)

type Pair struct {
	Row int
	Col int
}

type Map struct {
	Data  map[Pair]byte
	Start Pair
	End   Pair
}

func ReadInput(fname string) Map {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	m := Map{make(map[Pair]byte), Pair{}, Pair{}}
	i := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		for j, c := range sc.Text() {
			m.Data[Pair{i, j}] = byte(c)
			if c == 'S' {
				m.Start = Pair{i, j}
			}
			if c == 'E' {
				m.End = Pair{i, j}
			}
		}
		i++
	}
	return m
}

func GetBestTime(m Map, start Pair) int {
	paths := make(map[Pair]int)
	seen := make(map[Pair]bool)
	queue := list.New()
	queue.PushBack(start)
	seen[start] = true
	dirs := []Pair{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for queue.Len() > 0 {
		p := queue.Remove(queue.Front()).(Pair)
		if p == m.End {
			break
		}
		for _, d := range dirs {
			newP := Pair{p.Row + d.Row, p.Col + d.Col}
			val, ok := m.Data[newP]
			if !ok || val == '#' || seen[newP] {
				continue
			}
			seen[newP] = true
			paths[newP] = paths[p] + 1
			queue.PushBack(newP)
		}
	}
	return paths[m.End]
}

func GetBestFairPath(m Map) []Pair {
	paths := make(map[Pair]Pair)
	seen := make(map[Pair]bool)
	queue := list.New()
	queue.PushBack(m.Start)
	seen[m.Start] = true
	dirs := []Pair{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for queue.Len() > 0 {
		p := queue.Remove(queue.Front()).(Pair)
		if p == m.End {
			break
		}
		for _, d := range dirs {
			newP := Pair{p.Row + d.Row, p.Col + d.Col}
			val, ok := m.Data[newP]
			if !ok || val == '#' || seen[newP] {
				continue
			}
			seen[newP] = true
			paths[newP] = p
			queue.PushBack(newP)
		}
	}
	res := []Pair{m.End}
	for res[len(res)-1] != m.Start {
		res = append(res, paths[res[len(res)-1]])
	}
	slices.Reverse(res)
	return res
}

func CalculateDirs(t int) []Pair {
	if t <= 0 {
		return []Pair{{0, 0}}
	}
	result := CalculateDirs(t - 1)
	seen := make(map[Pair]bool)
	queue := list.New()
	for _, p := range result {
		seen[p] = true
		queue.PushBack(p)
	}
	dirs := []Pair{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for queue.Len() > 0 {
		p := queue.Remove(queue.Front()).(Pair)
		for _, d := range dirs {
			newP := Pair{p.Row + d.Row, p.Col + d.Col}
			if utils.AbsInt(newP.Row) > t || utils.AbsInt(newP.Col) > t || seen[newP] {
				continue
			}
			seen[newP] = true
			if utils.AbsInt(newP.Row)+utils.AbsInt(newP.Col) == t {
				result = append(result, newP)
			}
			queue.PushBack(newP)
		}
	}
	return result
}

func CountWithCheats(m Map, init []Pair, save, cheatT int) int {
	dirs := CalculateDirs(cheatT)
	result := 0
	for t, p := range init {
		for _, d := range dirs {
			newP := Pair{p.Row + d.Row, p.Col + d.Col}
			val, ok := m.Data[newP]
			if !ok || val == '#' {
				continue
			}
			idx := slices.Index(init, newP)
			var newT int
			curCheatT := utils.AbsInt(d.Row) + utils.AbsInt(d.Col)
			if idx != -1 {
				newT = t + curCheatT + len(init[idx:])
			} else {
				newT = t + curCheatT + GetBestTime(m, newP)
			}
			if newT < len(init)-1 && len(init)-newT >= save {
				result += 1
			}
		}
	}
	return result
}

func SolvePartOne(fname string) {
	m := ReadInput(fname)
	path := GetBestFairPath(m)
	fmt.Println(CountWithCheats(m, path, 100, 2))
}

func SolvePartTwo(fname string) {
	m := ReadInput(fname)
	path := GetBestFairPath(m)
	fmt.Println(CountWithCheats(m, path, 100, 20))
}
