package day12

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Pair struct {
	Row int
	Col int
}

type Region struct {
	Data map[Pair]rune
}

func (r *Region) Area() int {
	return len(r.Data)
}

func (r *Region) Perimeter() int {
	dirs := []Pair{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	result := 0
	for p := range r.Data {
		for _, dir := range dirs {
			newP := Pair{p.Row + dir.Row, p.Col + dir.Col}
			_, ok := r.Data[newP]
			if !ok {
				result += 1
			}
		}
	}
	return result
}

func (r *Region) Contains(p Pair) bool {
	_, ok := r.Data[p]
	return ok
}

func (r *Region) Sides() int {
	result := 0
	mn, mx := r.MinMax()
	// check rows
	for _, dir := range []Pair{{-1, 0}, {1, 0}} {
		for i := mn.Row; i < mx.Row+1; i++ {
			corner := false
			for j := mn.Col; j < mx.Col+1; j++ {
				p := Pair{i, j}
				if !r.Contains(p) {
					if corner {
						corner = false
						result += 1
					}
					continue
				}

				sideP := Pair{p.Row + dir.Row, p.Col + dir.Col}
				if r.Contains(sideP) {
					if corner {
						corner = false
						result += 1
					}
				} else {
					corner = true
				}
			}
			if corner {
				result += 1
			}
		}
	}
	// check cols
	for _, dir := range []Pair{{0, -1}, {0, 1}} {
		for j := mn.Col; j < mx.Col+1; j++ {
			corner := false
			for i := mn.Row; i < mx.Row+1; i++ {
				p := Pair{i, j}
				if !r.Contains(p) {
					if corner {
						corner = false
						result += 1
					}
					continue
				}

				sideP := Pair{p.Row + dir.Row, p.Col + dir.Col}
				if r.Contains(sideP) {
					if corner {
						corner = false
						result += 1
					}
				} else {
					corner = true
				}
			}
			if corner {
				result += 1
			}
		}
	}
	return result
}

func (r *Region) MinMax() (minRowCol, maxRowCol Pair) {
	minRowCol = Pair{999999, 999999}
	maxRowCol = Pair{-999999, -999999}
	for p := range r.Data {
		minRowCol.Row = min(minRowCol.Row, p.Row)
		minRowCol.Col = min(minRowCol.Col, p.Col)
		maxRowCol.Row = max(maxRowCol.Row, p.Row)
		maxRowCol.Col = max(maxRowCol.Col, p.Col)
	}
	return
}

func ReadInput(fname string) map[Pair]rune {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	res := make(map[Pair]rune)
	sc := bufio.NewScanner(file)
	i := 0
	for sc.Scan() {
		for j, c := range sc.Text() {
			res[Pair{i, j}] = c
		}
		i++
	}
	return res
}

func ScanRegion(data map[Pair]rune, seen map[Pair]bool, start Pair) Region {
	dirs := []Pair{
		{-1, 0}, {0, 1}, {1, 0}, {0, -1},
	}
	queue := list.New()
	queue.PushBack(start)
	r := Region{make(map[Pair]rune)}
	for queue.Len() != 0 {
		p := queue.Remove(queue.Front()).(Pair)
		seen[p] = true
		if data[p] == data[start] {
			r.Data[p] = data[p]
		}
		for _, dir := range dirs {
			newP := Pair{p.Row + dir.Row, p.Col + dir.Col}
			val, ok := data[newP]
			if !ok || val != data[start] || seen[newP] {
				continue
			}
			queue.PushBack(newP)
		}
	}
	return r
}

func DivideToRegions(data map[Pair]rune) []Region {
	regions := make([]Region, 0)
	seen := make(map[Pair]bool)
	for p := range data {
		if seen[p] {
			continue
		}
		regions = append(regions, ScanRegion(data, seen, p))
	}
	return regions
}

func CalculatePricesWithSides(data map[Pair]rune) int {
	seen := make(map[Pair]bool)
	total := 0
	for p := range data {
		if seen[p] {
			continue
		}
		region := ScanRegion(data, seen, p)
		sides := region.Sides()
		total += region.Area() * sides
	}
	return total
}

func SolvePartOne(fname string) {
	data := ReadInput(fname)
	regions := DivideToRegions(data)
	total := 0
	for _, r := range regions {
		total += r.Area() * r.Perimeter()
	}
	fmt.Println(total)
}

func SolvePartTwo(fname string) {
	data := ReadInput(fname)
	fmt.Println(CalculatePricesWithSides(data))
}
