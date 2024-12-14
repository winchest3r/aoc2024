package day14

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/winchest3r/aoc2024/utils"
)

type Pair struct {
	X, Y int
}

type Robot struct {
	Pos Pair
	Vel Pair
}

type Space struct {
	Robots []Robot
	Limits Pair
}

func (s *Space) Print() {
	data := make([][]byte, s.Limits.Y)
	for i := 0; i < s.Limits.Y; i++ {
		for j := 0; j < s.Limits.X; j++ {
			data[i] = append(data[i], '.')
		}
	}
	for _, r := range s.Robots {
		data[r.Pos.Y][r.Pos.X] = 'x'
	}
	for _, row := range data {
		fmt.Println(string(row))
	}
}

func (s *Space) CountRobotsInRadius(data map[Pair]bool, seen map[Pair]bool, cur Pair, dirs []Pair, rad int) int {
	result := 0
	if data[cur] {
		result = 1
	}
	if rad <= 0 {
		return result
	}
	for _, d := range dirs {
		newCur := Pair{cur.Y + d.Y, cur.X + d.X}
		if !seen[newCur] {
			seen[newCur] = true
			result += s.CountRobotsInRadius(data, seen, newCur, dirs, rad-1)
		}
	}
	return result
}

// HaveSomeDensity checks if there is some unusual pattern in space
func (s *Space) HaveSomeDensity(rad, closest, count int) bool {
	dirs := []Pair{
		{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1},
	}
	data := make(map[Pair]bool)
	for _, r := range s.Robots {
		data[r.Pos] = true
	}
	counter := 0
	for _, r := range s.Robots {
		seen := make(map[Pair]bool)
		seen[r.Pos] = true
		check := s.CountRobotsInRadius(data, seen, r.Pos, dirs, rad)
		if check > count {
			counter++
		}
	}
	return counter > len(s.Robots)/8
}

func (s *Space) SafetyFactor() int {
	count := []int{0, 0, 0, 0}
	for _, r := range s.Robots {
		mid := Pair{s.Limits.X / 2, s.Limits.Y / 2}
		if r.Pos.X == mid.X || r.Pos.Y == mid.Y {
			continue
		}
		quadrant := -1
		switch {
		case r.Pos.Y < mid.Y:
			if r.Pos.X < mid.X {
				quadrant = 0
			} else {
				quadrant = 1
			}
		case r.Pos.Y > mid.Y:
			if r.Pos.X < mid.X {
				quadrant = 2
			} else {
				quadrant = 3
			}
		}
		count[quadrant]++
	}
	result := 1
	for _, val := range count {
		result *= val
	}
	return result
}

func ReadInput(fname string) ([]Pair, []Pair) {
	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pattern := regexp.MustCompile(`^p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)$`)
	positions := make([]Pair, 0)
	velocities := make([]Pair, 0)
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		m := pattern.FindStringSubmatch(sc.Text())
		vals := utils.CastToInt(m[1:])
		positions = append(positions, Pair{vals[0], vals[1]})
		velocities = append(velocities, Pair{vals[2], vals[3]})
	}
	return positions, velocities
}

func CreateSpace(poses, vels []Pair, limits Pair) Space {
	space := Space{
		Robots: make([]Robot, 0, len(poses)),
		Limits: limits,
	}
	for i := range poses {
		space.Robots = append(space.Robots, Robot{poses[i], vels[i]})
	}
	return space
}

func ProcessSpace(s Space, seconds int) {
	for i := range s.Robots {
		x := (s.Robots[i].Pos.X + (s.Limits.X+s.Robots[i].Vel.X)%s.Limits.X*seconds) % s.Limits.X
		y := (s.Robots[i].Pos.Y + (s.Limits.Y+s.Robots[i].Vel.Y)%s.Limits.Y*seconds) % s.Limits.Y
		s.Robots[i].Pos = Pair{x, y}
	}
}

func ScanSpaceAndTryToFindChristmasTree(s Space, seconds int) {
	rad, one, count := 5, 4, 20
	for i := 0; i < seconds; i++ {
		ProcessSpace(s, 1)
		if s.HaveSomeDensity(rad, one, count) {
			fmt.Printf("Found some big density in space at %d second.\n", i+1)
			s.Print()
			fmt.Println("--------------------------------------------------")
		}
	}
}

func SolvePartOne(fname string) {
	p, v := ReadInput(fname)
	s := CreateSpace(p, v, Pair{101, 103})
	ProcessSpace(s, 100)
	fmt.Println(s.SafetyFactor())
}

func SolvePartTwo(fname string) {
	p, v := ReadInput(fname)
	s := CreateSpace(p, v, Pair{101, 103})
	ScanSpaceAndTryToFindChristmasTree(s, 10000)
}
