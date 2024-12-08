package day08

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	Row int
	Col int
}

func ReadInput(fname string) [][]byte {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make([][]byte, 0)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		data = append(data, []byte(line))
	}
	return data
}

func GetAntennas(data [][]byte) map[byte][]Pair {
	result := make(map[byte][]Pair)
	for i, row := range data {
		for j, freq := range row {
			if freq == '.' {
				continue
			}
			result[freq] = append(result[freq], Pair{i, j})
		}
	}
	return result
}

func CalculateAntinodes(data [][]byte) map[Pair]bool {
	antMap := GetAntennas(data)
	result := make(map[Pair]bool)
	for freq := range antMap {
		a := antMap[freq]
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a); j++ {
				if i == j {
					continue
				}
				diff := Pair{a[j].Row - a[i].Row, a[j].Col - a[i].Col}
				aNode := Pair{a[j].Row + diff.Row, a[j].Col + diff.Col}
				if aNode.Row < 0 || aNode.Row >= len(data) || aNode.Col < 0 || aNode.Col >= len(data[0]) {
					continue
				}
				result[aNode] = true
			}
		}
	}
	return result
}

func CalculateResonantAntinodes(data [][]byte) map[Pair]bool {
	antMap := GetAntennas(data)
	result := make(map[Pair]bool)
	for freq := range antMap {
		a := antMap[freq]
		for i := 0; i < len(a); i++ {
			for j := 0; j < len(a); j++ {
				if i == j {
					result[a[i]] = true
					continue
				}
				diff := Pair{a[j].Row - a[i].Row, a[j].Col - a[i].Col}
				aNode := Pair{a[j].Row + diff.Row, a[j].Col + diff.Col}
				for aNode.Row >= 0 && aNode.Row < len(data) && aNode.Col >= 0 && aNode.Col < len(data[0]) {
					result[aNode] = true
					aNode = Pair{aNode.Row + diff.Row, aNode.Col + diff.Col}
				}
			}
		}
	}
	return result
}

func SolvePartOne(fname string) {
	result := CalculateAntinodes(ReadInput(fname))
	fmt.Println(len(result))
}

func SolvePartTwo(fname string) {
	result := CalculateResonantAntinodes(ReadInput(fname))
	fmt.Println(len(result))
}
