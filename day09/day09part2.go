package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Block struct {
	Data     []BlockUnit
	Size     int
	InitSize int
}

type MemoryTwo struct {
	Data []Block
}

func NewMemoryTwo() *MemoryTwo {
	return &MemoryTwo{
		make([]Block, 0, 256),
	}
}

func ReadInputPartTwo(fname string) *MemoryTwo {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		data = sc.Text()
	}

	mem := NewMemoryTwo()
	for i, c := range data {
		sz, err := strconv.Atoi(string(c))
		if sz == 0 {
			continue
		}
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			mem.Data = append(mem.Data, Block{make([]BlockUnit, 0, sz), sz, sz})
			for j := 0; j < sz; j++ {
				mem.Data[len(mem.Data)-1].Data = append(mem.Data[len(mem.Data)-1].Data, BlockUnit{i / 2})
			}
		} else {
			for j := 0; j < sz; j++ {
				mem.Data[len(mem.Data)-1].Data = append(mem.Data[len(mem.Data)-1].Data, BlockUnit{-1})
			}
		}
	}
	return mem
}

func (m *MemoryTwo) FindFree(sz, limit int) int {
	for i, b := range m.Data[:limit] {
		if len(b.Data)-b.Size >= sz {
			return i
		}
	}
	return -1
}

func FragmentDiskTwo(m *MemoryTwo) {
	for i := len(m.Data) - 1; i > 0; i-- {
		free := m.FindFree(m.Data[i].InitSize, i)
		if free != -1 {
			k := 0
			for j, b := range m.Data[free].Data {
				if m.Data[i].Size == 0 {
					break
				}
				if b.Id == -1 {
					m.Data[free].Data[j].Id = m.Data[i].Data[k].Id
					m.Data[free].Size++
					m.Data[i].Data[k].Id = -2
					k++
					m.Data[i].Size--
				}
			}
		}
	}
}

func CalculateChecksumTwo(m *MemoryTwo) int {
	result := 0
	i := 0
	for _, b := range m.Data {
		for _, u := range b.Data {
			if u.Id >= 0 {
				result += i * u.Id
			}
			i++
		}
	}
	return result
}

func SolvePartTwo(fname string) {
	mem := ReadInputPartTwo(fname)
	FragmentDiskTwo(mem)
	fmt.Println(CalculateChecksumTwo(mem))
}
