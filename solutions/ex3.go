package solutions

import "fmt"

func Ex3a(lines []string) {
	count := 0
	le := len(lines[0])
	gMap := map[int]map[int]rune{}
	for i, line := range lines {
		for j, c := range line {
			if _, ok := gMap[j]; !ok {
				gMap[j] = map[int]rune{}
			}
			gMap[j][i] = c
		}
	}

	x := 3
	y := 1

	for i := 1; i < len(lines); i++ {
		if gMap[x][y] == '#' {
			count++
		}
		x = (x + 3) % le
		y++
	}

	fmt.Println(count)
}

func Ex3b(lines []string) {
	le := len(lines[0])
	gMap := map[int]map[int]rune{}
	for i, line := range lines {
		for j, c := range line {
			if _, ok := gMap[j]; !ok {
				gMap[j] = map[int]rune{}
			}
			gMap[j][i] = c
		}
	}

	xx := []int{1, 3, 5, 7, 1}
	x := []int{1, 3, 5, 7, 1}
	yy := []int{1, 1, 1, 1, 2}
	y := []int{1, 1, 1, 1, 2}
	cc := []int{0, 0, 0, 0, 0}

	for i := 1; i < len(lines); i++ {
		for j := range xx {
			if gMap[x[j]][y[j]] == '#' {
				cc[j] = cc[j] + 1
			}
			x[j] = (x[j] + xx[j]) % le
			y[j] = y[j] + yy[j]
		}
	}

	count := 1
	for _, v := range cc {
		count *= v
	}

	fmt.Println(count)
}
