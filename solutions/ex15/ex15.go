package ex15

import (
	"fmt"
	"strconv"
	"strings"
)

const Path = "./solutions/ex15"

func A(lines []string) {
	count := 0

	input := strings.Split(lines[0], ",")
	game := map[int]int{}
	spoken := [2021]int{}

	for i, n := range input {
		num, _ := strconv.ParseInt(n, 10, 64)
		game[int(num)] = i + 1
		spoken[i+1] = int(num)
	}

	for i := len(input) + 1; i <= 2020; i++ {
		lastBefore := 0
		for j := i - 2; j >= 0; j-- {
			if spoken[i-1] == spoken[j] {
				lastBefore = j
				break
			}
		}
		if lastBefore == 0 {
			spoken[i] = 0
		} else {
			spoken[i] = i - 1 - lastBefore
		}
	}

	count = spoken[2020]

	fmt.Println(count)
}

type pair struct {
	x, y int
}

func B(lines []string) {
	count := 0

	input := strings.Split(lines[0], ",")
	game := map[int]pair{}
	lastSpoken := 0

	for i, n := range input {
		num, _ := strconv.ParseInt(n, 10, 64)
		game[int(num)] = pair{i + 1, 0}
		lastSpoken = int(num)
	}

	for i := len(input) + 1; i <= 30000000; i++ {

		toSayPair := game[lastSpoken]
		if toSayPair.y == 0 {
			lastSpoken = 0
		} else {
			lastSpoken = toSayPair.x - toSayPair.y
		}
		p := game[lastSpoken]
		game[lastSpoken] = pair{i, p.x}

	}

	count = lastSpoken

	fmt.Println(count)
}
