package ex11

import (
	"fmt"
)

const Path = "./solutions/ex11"

type board struct {
	data  map[int]map[int]rune
	empty rune
}

func A(lines []string) {
	count := 0

	b := loadBoard(lines, '.')

	b, _ = iterate(b)

	for _, m := range b.data {
		for _, v := range m {
			if v == '#' {
				count++
			}
		}
	}

	fmt.Println(count)
}

func B(lines []string) {
	count := 0

	b := loadBoard(lines, 'e')

	b, _ = iterate2(b)

	for _, m := range b.data {
		for _, v := range m {
			if v == '#' {
				count++
			}
		}
	}

	fmt.Println(count)
}

func loadBoard(lines []string, empty rune) board {
	b := board{
		data:  map[int]map[int]rune{},
		empty: empty,
	}
	for y, l := range lines {
		for x, c := range l {
			if _, ok := b.data[x]; !ok {
				b.data[x] = map[int]rune{}
			}
			b.data[x][y] = c
		}
	}

	return b
}

func (b board) get(x, y int) rune {
	if _, ok := b.data[x]; ok {
		if v, ok2 := b.data[x][y]; ok2 {
			return v
		}
	}
	return b.empty
}

func deepcopy(b board) board {
	bc := board{
		data:  map[int]map[int]rune{},
		empty: b.empty,
	}
	for x, m := range b.data {
		for y, v := range m {
			if _, ok := bc.data[x]; !ok {
				bc.data[x] = map[int]rune{}
			}
			bc.data[x][y] = v
		}
	}
	return bc
}

func iterate(b board) (board, int) {
	count := 0
	for {
		prev := deepcopy(b)
		changed := 0

		for x, m := range prev.data {
			for y, _ := range m {
				if v := prev.get(x, y); v == 'L' {
					if prev.numAround(x, y) == 0 {
						b.data[x][y] = '#'
						changed++
					}
				} else if v == '#' {
					if prev.numAround(x, y) >= 4 {
						b.data[x][y] = 'L'
						changed++
					}
				}
			}
		}
		if changed == 0 {
			break
		}
		count++
	}

	return b, count
}

func (b board) numAround(x, y int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if b.get(x+dx, y+dy) == '#' {
				count++
			}
		}
	}
	return count
}

func iterate2(b board) (board, int) {
	count := 0
	for {
		prev := deepcopy(b)
		changed := 0

		for x, m := range prev.data {
			for y, _ := range m {
				if v := prev.get(x, y); v == 'L' {
					if prev.numAround2(x, y) == 0 {
						b.data[x][y] = '#'
						changed++
					}
				} else if v == '#' {
					if prev.numAround2(x, y) >= 5 {
						b.data[x][y] = 'L'
						changed++
					}
				}
			}
		}
		if changed == 0 {
			break
		}
		count++
	}

	return b, count
}

func (b board) numAround2(x, y int) int {
	count := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			xx := x + dx
			yy := y + dy
			for {
				v := b.get(xx, yy)
				if v == b.empty || v == 'L' {
					break
				}
				if v == '#' {
					count++
					break
				}
				xx += dx
				yy += dy
			}
		}
	}
	return count
}
