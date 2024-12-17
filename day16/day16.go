package day16

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type Pair struct {
	Row int
	Col int
}

func (a Pair) Diff(b Pair) Pair {
	a.Row -= b.Row
	a.Col -= b.Col
	return a
}

type Elem struct {
	Pair     Pair
	Prev     Pair
	Priority int
	index    int
}

type PriorityQueue []*Elem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(p any) {
	n := len(*pq)
	item := p.(*Elem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func ReadInput(fname string) (map[Pair]rune, Pair, Pair) {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make(map[Pair]rune)
	beg, end := Pair{}, Pair{}
	i := 0
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		for j, c := range sc.Text() {
			data[Pair{i, j}] = c
			if c == 'S' {
				beg = Pair{i, j}
			} else if c == 'E' {
				end = Pair{i, j}
			}
		}
		i++
	}
	return data, beg, end
}

func CalculateBestPath(data map[Pair]rune, beg, end Pair) int {
	dirs := []Pair{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
	seen := make(map[Pair]bool)
	weights := make(map[Pair]int)
	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Elem{beg, Pair{beg.Row, beg.Col - 1}, 0, 0})
	seen[beg] = true
	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*Elem)
		if p.Pair == end {
			break
		}
		for _, d := range dirs {
			newP := Pair{p.Pair.Row + d.Row, p.Pair.Col + d.Col}
			if seen[newP] || data[newP] == '#' || newP == p.Prev {
				continue
			}
			seen[newP] = true
			pr := 0
			if d != p.Pair.Diff(p.Prev) {
				weights[newP] = weights[p.Pair] + 1001
			} else {
				weights[newP] = weights[p.Pair] + 1
				pr = 1
			}
			heap.Push(&pq, &Elem{newP, p.Pair, pr, 0})
		}
	}
	return weights[end]
}

func SolvePartOne(fname string) {
	val := CalculateBestPath(ReadInput(fname))
	fmt.Println(val)
}

func SolvePartTwo(fname string) {

}
