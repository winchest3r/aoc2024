package day04

import (
	"bufio"
	"fmt"
	"os"
)

type Pair struct {
	I int
	J int
}

type CharPair struct {
	C byte
	P Pair
}

func CheckWord(data [][]byte, word string, c, i, j int, d Pair) bool {
	// we found the word
	if c >= len(word) {
		return true
	}
	rows, cols := len(data), len(data[0])
	// out of range or wrong char
	if i < 0 || j < 0 || i >= rows || j >= cols || data[i][j] != word[c] {
		return false
	}

	return CheckWord(data, word, c+1, i+d.I, j+d.J, d)
}

func MasChar(data [][]byte, c byte, i, j int) bool {
	if i < 0 || j < 0 || i >= len(data) || j >= len(data[0]) {
		return false
	}

	return c == data[i][j]
}

func XmasCounter(data [][]byte) int {
	dirs := []Pair{
		{-1, 0}, {-1, 1}, {0, 1},
		{1, 1}, {1, 0},
		{1, -1}, {0, -1}, {-1, -1},
	}

	words := 0
	for i, row := range data {
		for j, c := range row {
			if c != 'X' {
				continue
			}
			// check all directions
			for _, d := range dirs {
				if CheckWord(data, "XMAS", 0, i, j, d) {
					words++
				}
			}
		}
	}
	return words
}

func MasCounter(data [][]byte) int {
	words := 0
	for i, row := range data {
		for j, c := range row {
			if c != 'A' {
				continue
			}

			iTopLeft, jTopLeft := i-1, j-1
			iTopRight, jTopRight := i-1, j+1
			iDownLeft, jDownLeft := i+1, j-1
			iDownRight, jDownRight := i+1, j+1

			// check diagonals
			leftDiagonal := MasChar(data, 'M', iTopLeft, jTopLeft) && MasChar(data, 'S', iDownRight, jDownRight) ||
				MasChar(data, 'S', iTopLeft, jTopLeft) && MasChar(data, 'M', iDownRight, jDownRight)

			rightDiagonal := MasChar(data, 'M', iTopRight, jTopRight) && MasChar(data, 'S', iDownLeft, jDownLeft) ||
				MasChar(data, 'S', iTopRight, jTopRight) && MasChar(data, 'M', iDownLeft, jDownLeft)

			if leftDiagonal && rightDiagonal {
				words++
			}
		}
	}
	return words
}

func ReadInput(fname string) [][]byte {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	result := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			result = append(result, []byte(text))
		}
	}
	return result
}

func SolvePartOne(fname string) {
	data := ReadInput(fname)
	result := XmasCounter(data)
	fmt.Println(result)
}

func SolvePartTwo(fname string) {
	data := ReadInput(fname)
	result := MasCounter(data)
	fmt.Println(result)
}
