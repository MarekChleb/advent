package ex17

import (
	"fmt"
)

const Path = "./solutions/ex17"

type triplet struct {
	x, y, z int
}

type grid struct {
	data       map[triplet]int
	potentials map[triplet]int
}

type quadriplet struct {
	x, y, z, i int
}

type hyperGrid struct {
	data       map[quadriplet]int
	potentials map[quadriplet]int
}

func A(lines []string) {
	count := 0

	g := newGrid()

	for y, l := range lines {
		for x, c := range l {
			if c == '#' {
				g.insert(x, y, 0)
			}
		}
	}

	for i := 0; i < 6; i++ {
		g.ascend()
		fmt.Printf("test: %#v\n", g.count())
	}

	count = g.count()

	fmt.Println(count)
}

func newGrid() grid {
	return grid{map[triplet]int{}, map[triplet]int{}}
}

func (g *grid) insert(x, y, z int) {
	g.data[triplet{x, y, z}] = 1
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dx == 0 && dz == 0 && dy == 0 {
					continue
				}
				xyz := triplet{x + dx, y + dy, z + dz}
				g.potentials[xyz] = g.potentials[xyz] + 1
			}
		}
	}
}

func (g *grid) ascend() {
	newG := newGrid()

	for t, p := range g.potentials {
		if p == 3 {
			newG.insert(t.x, t.y, t.z)
			continue
		}
		if p == 2 && g.data[t] == 1 {
			newG.insert(t.x, t.y, t.z)
		}
	}

	*g = newG
}

func (g grid) count() int {
	count := 0
	for _, p := range g.data {
		if p == 1 {
			count++
		}
	}
	return count
}

func B(lines []string) {
	count := 0

	g := newHyperGrid()

	for y, l := range lines {
		for x, c := range l {
			if c == '#' {
				g.insert(x, y, 0, 0)
			}
		}
	}

	for i := 0; i < 6; i++ {
		g.ascend()
		fmt.Printf("test: %#v\n", g.count())
	}

	count = g.count()

	fmt.Println(count)
}

func newHyperGrid() hyperGrid {
	return hyperGrid{map[quadriplet]int{}, map[quadriplet]int{}}
}

func (g *hyperGrid) insert(x, y, z, i int) {
	g.data[quadriplet{x, y, z, i}] = 1
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for di := -1; di <= 1; di++ {
					if dx == 0 && dz == 0 && dy == 0 && di == 0 {
						continue
					}
					xyzi := quadriplet{x + dx, y + dy, z + dz, i + di}
					g.potentials[xyzi] = g.potentials[xyzi] + 1
				}
			}
		}
	}
}

func (g *hyperGrid) ascend() {
	newG := newHyperGrid()

	for t, p := range g.potentials {
		if p == 3 {
			newG.insert(t.x, t.y, t.z, t.i)
			continue
		}
		if p == 2 && g.data[t] == 1 {
			newG.insert(t.x, t.y, t.z, t.i)
		}
	}

	*g = newG
}

func (g hyperGrid) count() int {
	count := 0
	for _, p := range g.data {
		if p == 1 {
			count++
		}
	}
	return count
}
