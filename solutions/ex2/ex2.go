package ex2

import (
	"fmt"
	"strconv"
	"strings"
)

type pass struct {
	min, max int
	letter   rune
	p        string
}

const Path = "./solutions/ex2"

func A(lines []string) {

	count := 0

	for _, l := range lines {
		lEls := strings.Split(l, " ")
		lNs := strings.Split(lEls[0], "-")
		min, _ := strconv.ParseInt(lNs[0], 10, 64)
		max, _ := strconv.ParseInt(lNs[1], 10, 64)

		var letter rune
		for i, v := range lEls[1] {
			if i == 0 {
				letter = v
			}
		}

		pp := pass{
			min:    int(min),
			max:    int(max),
			letter: letter,
			p:      lEls[2],
		}
		fmt.Println(pp)

		if pp.check() {
			count++
		}
	}

	fmt.Println(count)
}

func B(lines []string) {

	count := 0

	for _, l := range lines {
		lEls := strings.Split(l, " ")
		lNs := strings.Split(lEls[0], "-")
		min, _ := strconv.ParseInt(lNs[0], 10, 64)
		max, _ := strconv.ParseInt(lNs[1], 10, 64)

		var letter rune
		for i, v := range lEls[1] {
			if i == 0 {
				letter = v
			}
		}

		pp := pass{
			min:    int(min),
			max:    int(max),
			letter: letter,
			p:      lEls[2],
		}
		fmt.Println(pp)

		if pp.check2() {
			count++
		}
	}

	fmt.Println(count)
}

func (p pass) check() bool {
	count := 0
	for _, s := range p.p {
		if s == p.letter {
			count++
		}
	}
	return count >= p.min && count <= p.max
}

func (p pass) check2() bool {
	ok := false
	for i, s := range p.p {
		if i+1 == p.min {
			ok = s == p.letter
		}
		if i+1 == p.max {

			return ok != (s == p.letter)
		}
	}
	return false
}
