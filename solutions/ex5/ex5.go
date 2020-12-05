package ex5

import (
	"fmt"
	"math"
)

var translate = map[rune]int{
	'F': 0,
	'B': 1,
	'L': 0,
	'R': 1,
}

const Path = "./solutions/ex5"

func A(lines []string) {
	maxID := 0

	for _, l := range lines {
		sum := 0
		for i, r := range l {
			sum += translate[r] * int(math.Pow(2, float64(9-i)))
		}
		if sum > maxID {
			maxID = sum
		}
	}

	fmt.Println(maxID)
}

func B(lines []string) {
	maxID := 0
	exists := map[int]struct{}{}

	for _, l := range lines {
		sum := 0
		for i, r := range l {
			sum += translate[r] * int(math.Pow(2, float64(9-i)))
		}
		if sum > maxID {
			maxID = sum
		}
		exists[sum] = struct{}{}
	}

	for i := maxID; i >= 0; i-- {
		if _, ok := exists[i]; !ok {
			fmt.Println(i)
			return
		}
	}

}
