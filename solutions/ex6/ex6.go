package ex6

import (
	"fmt"
)

const Path = "./solutions/ex6"

func A(lines []string) {
	count := 0

	var m = map[rune]struct{}{}

	for _, l := range lines {
		if l == "" {
			for range m {
				count++
			}
			m = map[rune]struct{}{}
			continue
		}

		for _, c := range l {
			m[c] = struct{}{}
		}
	}

	fmt.Println(count)
}

func B(lines []string) {
	count := 0

	var m = map[rune]int{}
	cc := 0

	for _, l := range lines {
		if l == "" {
			for _, v := range m {
				if v == cc {
					count++
				}
			}
			m = map[rune]int{}
			cc = 0
			continue
		}

		for _, c := range l {
			m[c] = m[c] + 1
		}
		cc++
	}

	fmt.Println(count)
}
