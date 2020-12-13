package ex12

import (
	"fmt"
	"regexp"
	"strconv"
)

const Path = "./solutions/ex12"

type ship struct {
	direction direction
	x, y      int
}

type direction struct {
	dx, dy int
}

func (d *direction) turnRight(val int) {
	num := val / 90
	for i := 0; i < num; i++ {
		d.dx, d.dy = d.dy, -d.dx
	}
}

func (d *direction) turnLeft(val int) {
	num := val / 90
	for i := 0; i < num; i++ {
		d.dx, d.dy = -d.dy, d.dx
	}
}

var directions = map[string]direction{
	"N": direction{0, 1},
	"S": direction{0, -1},
	"W": direction{-1, 0},
	"E": direction{1, 0},
}

func isDirection(c string) bool {
	if _, ok := directions[c]; ok {
		return true
	}
	return false
}

func A(lines []string) {
	count := 0

	s := ship{
		direction: direction{1, 0},
	}
	for _, l := range lines {
		s.action(l)
	}

	count = s.mDist()

	fmt.Println(count)
}

func B(lines []string) {
	count := 0

	s := ship{
		direction: direction{10, 1},
	}
	for _, l := range lines {
		s.actionRelative(l)
	}

	count = s.mDist()

	fmt.Println(count)
}

func (s *ship) action(line string) {
	mType, value := parseMove(line)

	if isDirection(mType) {
		dir := directions[mType]
		s.move(dir, value)
	}
	if mType == "F" {
		s.move(s.direction, value)
	}
	if mType == "L" {
		s.turnLeft(value)
	}
	if mType == "R" {
		s.turnRight(value)
	}
}

func (s *ship) actionRelative(line string) {
	mType, value := parseMove(line)

	if isDirection(mType) {
		dir := directions[mType]
		s.moveWaypoint(dir, value)
	}
	if mType == "F" {
		s.move(s.direction, value)
	}
	if mType == "L" {
		s.turnLeft(value)
	}
	if mType == "R" {
		s.turnRight(value)
	}
	fmt.Println(s.x, s.y)
}

func (s *ship) move(dir direction, value int) {
	s.x, s.y = s.x+dir.dx*value, s.y+dir.dy*value
}

func (s *ship) moveWaypoint(dir direction, value int) {
	s.direction.dx, s.direction.dy = s.direction.dx+dir.dx*value, s.direction.dy+dir.dy*value
}

func (s *ship) turnLeft(value int) {
	s.direction.turnLeft(value)
}

func (s *ship) turnRight(value int) {
	s.direction.turnRight(value)
}

func (s ship) mDist() int {
	x := s.x
	if x < 0 {
		x = -x
	}
	y := s.y
	if y < 0 {
		y = -y
	}
	return x + y
}

func parseMove(line string) (string, int) {
	re := regexp.MustCompile(`(N|S|E|W|F|R|L|F)([0-9]+)`)

	info := re.FindStringSubmatch(line)

	mType := info[1]
	valString := info[2]

	value, _ := strconv.ParseInt(valString, 10, 64)

	return mType, int(value)
}
