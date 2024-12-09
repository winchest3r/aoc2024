package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type BlockUnit struct {
	Id int
}

type Memory struct {
	Data []BlockUnit
}

func NewMemory() *Memory {
	return &Memory{
		make([]BlockUnit, 0, 256),
	}
}

func (m *Memory) AddBlockPartOne(id, size int) {
	for i := 0; i < size; i++ {
		m.Data = append(m.Data, BlockUnit{id})
	}
}

func (m *Memory) FindFree(idx int) int {
	for idx < len(m.Data) && m.Data[idx].Id != -1 {
		idx++
	}
	if idx >= len(m.Data) {
		return -1
	}
	return idx
}

func ReadInput(fname string) *Memory {
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

	mem := NewMemory()
	for i, c := range data {
		sz, err := strconv.Atoi(string(c))
		if err != nil {
			panic(err)
		}
		if i%2 == 0 {
			mem.AddBlockPartOne(i/2, sz)
		} else {
			mem.AddBlockPartOne(-1, sz)
		}
	}
	return mem
}

func FragmentDisk(m *Memory) {
	free := m.FindFree(0)
	idx := len(m.Data) - 1
	for free < idx {
		if m.Data[idx].Id != -1 {
			m.Data[free] = m.Data[idx]
			m.Data[idx].Id = -1
			free = m.FindFree(free + 1)
		}
		idx--
	}
}

func CalculateChecksum(m *Memory) int {
	result := 0
	for i, f := range m.Data {
		if f.Id != -1 {
			result += i * f.Id
		}
	}
	return result
}

func SolvePartOne(fname string) {
	mem := ReadInput(fname)
	FragmentDisk(mem)
	fmt.Println(CalculateChecksum(mem))
}

func SolvePartTwo(fname string) {

}
