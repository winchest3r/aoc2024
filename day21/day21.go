package day21

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Button struct {
	Value                 string
	Left, Right, Up, Down *Button
}

type Panel struct {
	Buttons []*Button
}

func NewDirectKeypad() Panel {
	bA := &Button{Value: "A"}
	bL := &Button{Value: "<"}
	bR := &Button{Value: ">"}
	bU := &Button{Value: "^"}
	bD := &Button{Value: "v"}

	/*
	       +---+---+
	       | ^ | A |
	   +---+---+---+
	   | < | v | > |
	   +---+---+---+
	*/

	bA.Left = bU
	bA.Down = bR

	bL.Right = bD

	bR.Up = bA
	bR.Left = bD

	bU.Right = bA
	bU.Down = bD

	bD.Up = bU
	bD.Right = bR
	bD.Left = bL

	return Panel{[]*Button{bA, bL, bR, bU, bD}}
}

func NewNumberKeypad() Panel {
	bA := &Button{Value: "A"}
	b0 := &Button{Value: "0"}
	b1 := &Button{Value: "1"}
	b2 := &Button{Value: "2"}
	b3 := &Button{Value: "3"}
	b4 := &Button{Value: "4"}
	b5 := &Button{Value: "5"}
	b6 := &Button{Value: "6"}
	b7 := &Button{Value: "7"}
	b8 := &Button{Value: "8"}
	b9 := &Button{Value: "9"}

	/*
	   +---+---+---+
	   | 7 | 8 | 9 |
	   +---+---+---+
	   | 4 | 5 | 6 |
	   +---+---+---+
	   | 1 | 2 | 3 |
	   +---+---+---+
	       | 0 | A |
	       +---+---+
	*/

	bA.Up = b3
	bA.Left = b0

	b0.Up = b2
	b0.Right = bA

	b1.Up = b4
	b1.Right = b2

	b2.Up = b5
	b2.Right = b3
	b2.Down = b0
	b2.Left = b1

	b3.Up = b6
	b3.Down = bA
	b3.Left = b2

	b4.Up = b7
	b4.Right = b5
	b4.Down = b1

	b5.Up = b8
	b5.Right = b6
	b5.Down = b2
	b5.Left = b4

	b6.Up = b9
	b6.Down = b3
	b6.Left = b5

	b7.Right = b8
	b7.Down = b4

	b8.Right = b9
	b8.Down = b5
	b8.Left = b7

	b9.Down = b6
	b9.Left = b8

	return Panel{[]*Button{bA, b0, b1, b2, b3, b4, b5, b6, b7, b8, b9}}
}

type Robot struct {
	Keypad Panel
	Pos    *Button
}

func findPaths(cur, end *Button, seen map[*Button]bool, path []*Button, allPaths *[][]*Button) {
	if cur == nil || seen[cur] {
		return
	}

	path = append(path, cur)

	if cur == end {
		*allPaths = append(*allPaths, path)
		return
	}

	seen[cur] = true

	findPaths(cur.Up, end, seen, path, allPaths)
	findPaths(cur.Right, end, seen, path, allPaths)
	findPaths(cur.Down, end, seen, path, allPaths)
	findPaths(cur.Left, end, seen, path, allPaths)

	seen[cur] = false
}

func FormPath(path []*Button) string {
	result := ""
	if len(path) == 0 {
		return result
	}
	start := path[0]
	for _, b := range path[1:] {
		switch b {
		case start.Up:
			result += "^"
		case start.Right:
			result += ">"
		case start.Down:
			result += "v"
		case start.Left:
			result += "<"
		}
		start = b
	}
	return result
}

func CountChanges(s string) int {
	if len(s) == 0 {
		return 0
	}
	changes := 0
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			changes++
		}
	}
	return changes
}

func FindBestPath(paths [][]*Button) string {
	bestPath, bestChanges := FormPath(paths[0]), CountChanges(FormPath(paths[0]))
	for _, p := range paths[1:] {
		newPath := FormPath(p)
		newChanges := CountChanges(newPath)
		if len(newPath) < len(bestPath) || newChanges < bestChanges {
			bestPath = newPath
			bestChanges = newChanges
		}
	}
	return bestPath
}

func (r *Robot) Move(pos string) string {
	allPaths := [][]*Button{}
	seen := make(map[*Button]bool)
	var end *Button
	for _, b := range r.Keypad.Buttons {
		if b.Value == pos {
			end = b
		}
	}
	findPaths(r.Pos, end, seen, []*Button{}, &allPaths)
	r.Pos = end
	return FindBestPath(allPaths)
}

func MakeSetOfMovements(r Robot, code string) string {
	result := ""
	for _, c := range code {
		move := r.Move(string(c))
		result += move + "A"
	}
	return result
}

func NewRobot(panel Panel, start string) Robot {
	r := Robot{}
	r.Keypad = panel
	for _, b := range r.Keypad.Buttons {
		if b.Value == start {
			r.Pos = b
			break
		}
	}
	return r
}

func ReadInput(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	codes := make([]string, 0)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		codes = append(codes, sc.Text())
	}
	return codes
}

func ThreeRobots(code string) string {
	t1 := NewNumberKeypad()
	t2 := NewDirectKeypad()
	t3 := NewDirectKeypad()
	r1 := NewRobot(t1, "A")
	r2 := NewRobot(t2, "A")
	r3 := NewRobot(t3, "A")
	for _, r := range []Robot{r1, r2, r3} {
		result := MakeSetOfMovements(r, code)
		code = result
	}
	return code
}

func ProcessRobots(robots []Robot, code string) string {
	for _, r := range robots {
		result := MakeSetOfMovements(r, code)
		code = result
	}
	return code
}

func CalculateSum(codes []string) int {
	result := 0
	for _, code := range codes {
		t1, t2, t3 := NewNumberKeypad(), NewDirectKeypad(), NewDirectKeypad()
		r1, r2, r3 := NewRobot(t1, "A"), NewRobot(t2, "A"), NewRobot(t3, "A")
		val, _ := strconv.Atoi(code[:len(code)-1])
		result += len(ProcessRobots([]Robot{r1, r2, r3}, code)) * val
	}
	return result
}

func SolvePartOne(fname string) {
	data := ReadInput(fname)
	fmt.Println(CalculateSum(data))
}

func SolvePartTwo(fname string) {

}
