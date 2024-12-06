package day06

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
)

type CellType int

const (
	Empty CellType = iota
	Visited
	Obstacle
	Invalid
)

type Coord struct {
	x int
	y int
}

type Dir struct {
	x int
	y int
}

func CheckCell(data map[Coord]CellType, pos Coord) CellType {
	val, ok := data[pos]
	if !ok {
		return Invalid
	}
	return val
}

// Move returns false if can't move
func Move(data map[Coord]CellType, pos *Coord, dir *Dir) bool {
	newPos := Coord{pos.x + dir.x, pos.y + dir.y}
	// turn while there is an obstacle
	for CheckCell(data, newPos) == Obstacle {
		*dir = NextDirection(*dir)
		newPos = Coord{pos.x + dir.x, pos.y + dir.y}
	}
	if CheckCell(data, newPos) == Invalid {
		return false
	}
	*pos = newPos
	data[*pos] = Visited
	return true
}

func GetVisitedCells(data map[Coord]CellType, pos Coord, dir Dir) int {
	for Move(data, &pos, &dir) {
	}
	counter := 0
	for key := range data {
		if data[key] == Visited {
			counter++
		}
	}
	return counter
}

func GetLoopsWithObstruction(data map[Coord]CellType, initPos Coord, initDir Dir) int {
	counter := 0
	for c := range data {
		if data[c] != Empty {
			continue
		}
		dataCopy := maps.Clone(data)
		// if we will have same coord with same direction - we in a loop
		visitedMap := make(map[Coord][]Dir)
		// make obstruction
		dataCopy[c] = Obstacle
		pos, dir := initPos, initDir
		visitedMap[c] = append(visitedMap[c], dir)
		for Move(dataCopy, &pos, &dir) {
			dirs, ok := visitedMap[pos]
			// if we visited this cell and it contains the same direction
			// leave it because it's a loop
			if ok && slices.Contains(dirs, dir) {
				counter++
				break
			}
			visitedMap[pos] = append(visitedMap[pos], dir)
		}
	}
	return counter
}

func GetDirection(c rune) Dir {
	switch c {
	case '^':
		return Dir{0, -1}
	case '>':
		return Dir{1, 0}
	case 'V':
		return Dir{0, 1}
	case '<':
		return Dir{-1, 0}
	}
	return Dir{}
}

func NextDirection(d Dir) Dir {
	switch d {
	case Dir{0, -1}:
		return Dir{1, 0}
	case Dir{1, 0}:
		return Dir{0, 1}
	case Dir{0, 1}:
		return Dir{-1, 0}
	case Dir{-1, 0}:
		return Dir{0, -1}
	}
	return Dir{}
}

func ReadInput(fname string) (map[Coord]CellType, Coord, Dir) {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make(map[Coord]CellType)
	var startCell Coord
	var startDir Dir
	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		x := 0
		for _, c := range row {
			switch c {
			case '.':
				data[Coord{x, y}] = Empty
			case '#':
				data[Coord{x, y}] = Obstacle
			default:
				startCell = Coord{x, y}
				data[startCell] = Visited
				startDir = GetDirection(c)
			}
			x++
		}
		y++
	}
	return data, startCell, startDir
}

func SolvePartOne(fname string) {
	result := GetVisitedCells(ReadInput(fname))
	fmt.Println(result)
}

func SolvePartTwo(fname string) {
	result := GetLoopsWithObstruction(ReadInput(fname))
	fmt.Println(result)
}
