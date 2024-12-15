package day15

import (
	"fmt"
	"os"
	"strings"
)

type Pair struct {
	Row int
	Col int
}

type Map struct {
	Data   map[Pair]byte
	Limits Pair
}

func (mp *Map) Robot() Pair {
	for p := range mp.Data {
		if mp.Data[p] == '@' {
			return p
		}
	}
	return Pair{}
}

func (mp *Map) Print() {
	for i := 0; i < mp.Limits.Row; i++ {
		for j := 0; j < mp.Limits.Col; j++ {
			fmt.Printf("%c", mp.Data[Pair{i, j}])
		}
		fmt.Println()
	}
}

func (p Pair) Mul(i int) Pair {
	p.Row *= i
	p.Col *= i
	return p
}

func (a Pair) Add(b Pair) Pair {
	a.Row += b.Row
	a.Col += b.Col
	return a
}

var Dirs map[byte]Pair = map[byte]Pair{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

func ReadInput(fname string) (*Map, []byte) {
	file, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\r\n\r\n")
	mp := &Map{Data: make(map[Pair]byte)}
	orders := make([]byte, 0)
	for i, row := range strings.Fields(data[0]) {
		for j, c := range row {
			mp.Data[Pair{i, j}] = byte(c)
			mp.Limits.Col = max(mp.Limits.Col, j)
		}
		mp.Limits.Row = max(mp.Limits.Row, i)
	}

	for _, row := range strings.Fields(data[1]) {
		for _, c := range row {
			orders = append(orders, byte(c))
		}
	}

	return mp, orders
}

func ReadInputPartTwo(fname string) (*Map, []byte) {
	file, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\r\n\r\n")
	mp := &Map{Data: make(map[Pair]byte)}
	orders := make([]byte, 0)
	for i, row := range strings.Fields(data[0]) {
		for j, c := range row {
			j1 := j * 2
			j2 := j*2 + 1
			switch c {
			case '#':
				mp.Data[Pair{i, j1}] = '#'
				mp.Data[Pair{i, j2}] = '#'
			case 'O':
				mp.Data[Pair{i, j1}] = '['
				mp.Data[Pair{i, j2}] = ']'
			case '.':
				mp.Data[Pair{i, j1}] = '.'
				mp.Data[Pair{i, j2}] = '.'
			case '@':
				mp.Data[Pair{i, j1}] = '@'
				mp.Data[Pair{i, j2}] = '.'
			}
			mp.Limits.Col = max(mp.Limits.Col, j*2+1)
		}
		mp.Limits.Row = max(mp.Limits.Row, i*2+1)
	}

	for _, row := range strings.Fields(data[1]) {
		for _, c := range row {
			orders = append(orders, byte(c))
		}
	}

	return mp, orders
}

func ProcessMove(mp *Map, prev, next, dir Pair) {
	if mp.Data[next] == '#' {
		return
	}
	if mp.Data[next] == '.' {
		mp.Data[next] = mp.Data[prev]
		mp.Data[prev] = '.'
		return
	}
	ProcessMove(mp, next, next.Add(dir), dir)
	if mp.Data[next] == '.' {
		mp.Data[next] = mp.Data[prev]
		mp.Data[prev] = '.'
	}
}

func Move(mp *Map, order byte) {
	d := Dirs[order]
	p := mp.Robot()
	ProcessMove(mp, p, p.Add(d), d)
}

func CalculateGPS(mp *Map) int {
	result := 0
	for p := range mp.Data {
		if mp.Data[p] == 'O' {
			gps := 100*p.Row + p.Col
			result += gps
		}
	}
	return result
}

func MoveInOrder(mp *Map, orders []byte) {
	for _, o := range orders {
		Move(mp, o)
	}
}

func SolvePartOne(fname string) {
	mp, orders := ReadInput(fname)
	MoveInOrder(mp, orders)
	fmt.Println(CalculateGPS(mp))
}

func SolvePartTwo(fname string) {

}
