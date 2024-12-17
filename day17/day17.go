package day17

import (
	"fmt"
	"math"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/winchest3r/aoc2024/utils"
)

type Machine struct {
	RegA    int
	RegB    int
	RegC    int
	Program []int
	Ptr     int
	Output  []int
}

func NewMachine(a, b, c int, ins []int) *Machine {
	return &Machine{a, b, c, ins, 0, make([]int, 0)}
}

func (m *Machine) Combo(operand int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return m.RegA
	case 5:
		return m.RegB
	case 6:
		return m.RegC
	}
	panic("couldn't get this far in ComboOperand")
}

func (m *Machine) Reset(a, b, c int) {
	m.RegA = a
	m.RegB = b
	m.RegC = c
	m.Ptr = 0
	m.Output = make([]int, 0)
}

func (m *Machine) Process() bool {
	if m.Ptr >= len(m.Program)-1 {
		return false
	}
	opcode := m.Program[m.Ptr]
	operand := m.Program[m.Ptr+1]
	m.Ptr += 2
	switch opcode {
	case 0:
		m.RegA /= int(math.Pow(2, float64(m.Combo(operand))))
	case 1:
		m.RegB ^= operand
	case 2:
		m.RegB = m.Combo(operand) % 8
	case 3:
		if m.RegA == 0 {
			return true
		}
		m.Ptr = operand
	case 4:
		m.RegB ^= m.RegC
	case 5:
		m.Output = append(m.Output, m.Combo(operand)%8)
	case 6:
		m.RegB = m.RegA / int(math.Pow(2, float64(m.Combo(operand))))
	case 7:
		m.RegC = m.RegA / int(math.Pow(2, float64(m.Combo(operand))))
	}
	return true
}

func (m *Machine) ProcessAll() {
	for m.Process() {
	}
}

func (m *Machine) OutputJoin() string {
	mapFunc := func(arr []int) []string {
		res := make([]string, 0, len(arr))
		for _, i := range arr {
			res = append(res, strconv.Itoa(i))
		}
		return res
	}
	return strings.Join(mapFunc(m.Output), ",")
}

func (m *Machine) OutputAndProgramIsEqual() bool {
	return reflect.DeepEqual(m.Output, m.Program)
}

func ReadInput(fname string) *Machine {
	file, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(file), "\r\n\r\n")
	p, err := regexp.Compile(`[^0-9]*(\d+)[^0-9]*(\d+)\s[^0-9]*(\d+)`)
	if err != nil {
		panic(err)
	}
	match := p.FindStringSubmatch(data[0])
	regs := utils.CastToInt(match[1:])

	ins := utils.SplitToInt(strings.Fields(data[1])[1], ",")
	machine := NewMachine(regs[0], regs[1], regs[2], ins)

	return machine
}

func SolvePartOne(fname string) {
	m := ReadInput(fname)
	m.ProcessAll()
	fmt.Println(m.OutputJoin())
}

func Factor8(i int) int {
	return int(math.Pow(8, float64(i)))
}

func CalculateA(f []int) int {
	res := 0
	for i, val := range f {
		res += Factor8(i) * val
	}
	return res
}

func FindCopy(m *Machine) int {
	f := make([]int, len(m.Program))
	for {
		a := CalculateA(f)
		m.Reset(a, 0, 0)
		m.ProcessAll()
		if m.OutputAndProgramIsEqual() {
			return a
		}
		for i := len(m.Program) - 1; i >= 0; i-- {
			if len(m.Output) < i || m.Output[i] != m.Program[i] {
				f[i] += 1
				break
			}
		}
	}
}

func SolvePartTwo(fname string) {
	m := ReadInput(fname)
	fmt.Println(FindCopy(m))
}
