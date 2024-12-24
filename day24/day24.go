package day24

import (
	"container/list"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Gate struct {
	Name  string
	Value byte
}

type Gates map[string]byte

func (g Gates) Clone() Gates {
	clone := make(Gates)
	for key, value := range g {
		clone[key] = value
	}
	return clone
}

type Operation struct {
	A, B, C string
	Op      func(byte, byte) byte
}

func And(a, b byte) byte {
	return a & b
}

func Or(a, b byte) byte {
	return a | b
}

func Xor(a, b byte) byte {
	return a ^ b
}

func ReadInput(fname string) (Gates, []Operation) {
	file, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(file), "\r\n\r\n")
	gates := strings.Fields(data[0])
	device := make(Gates)
	for i := 0; i < len(gates); i += 2 {
		name := gates[i][:len(gates[i])-1]
		value, _ := strconv.Atoi(gates[i+1])
		device[name] = byte(value)
	}
	ops := make([]Operation, 0)
	for _, line := range strings.Split(data[1], "\r\n") {
		if line == "" {
			continue
		}
		vals := strings.Fields(line)
		op := Operation{}
		op.A = vals[0]
		op.B = vals[2]
		op.C = vals[4]
		switch vals[1] {
		case "AND":
			op.Op = And
		case "OR":
			op.Op = Or
		case "XOR":
			op.Op = Xor
		}
		ops = append(ops, op)
	}
	return device, ops
}

func Process(gates Gates, ops []Operation) {
	opsList := list.New()
	for _, op := range ops {
		opsList.PushBack(op)
	}
	for opsList.Len() > 0 {
		op := opsList.Remove(opsList.Front()).(Operation)
		valA, okA := gates[op.A]
		valB, okB := gates[op.B]
		if okA && okB {
			gates[op.C] = op.Op(valA, valB)
		} else {
			opsList.PushBack(op)
		}
	}
}

func GetCombinedNumber(gates Gates, prefix string) int {
	data := make([]Gate, 0)
	for key := range gates {
		if strings.HasPrefix(key, prefix) {
			data = append(data, Gate{key, gates[key]})
		}
	}
	sort.Slice(data, func(i, j int) bool { return data[i].Name > data[j].Name })
	result := ""
	for _, g := range data {
		result += strconv.Itoa(int(g.Value))
	}
	i, _ := strconv.ParseInt(result, 2, 64)
	return int(i)
}

func SolvePartOne(fname string) {
	gates, ops := ReadInput(fname)
	Process(gates, ops)
	fmt.Println(GetCombinedNumber(gates, "z"))
}

func SolvePartTwo(fname string) {

}
