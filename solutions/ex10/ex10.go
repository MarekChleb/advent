package ex10

import (
	"fmt"
	"sort"
	"strconv"
)

const Path = "./solutions/ex10"

func A(lines []string) {
	count := 0

	jolts := []int{}

	for _, l := range lines {
		n, _ := strconv.ParseInt(l, 10, 64)
		jolts = append(jolts, int(n))
	}

	sort.Ints(jolts)

	ones := 0
	threes := 1
	prev := 0
	for _, n := range jolts {
		if n-prev == 1 {
			ones++
		} else if n-prev == 3 {
			threes++
		}
		prev = n
	}

	count = ones * threes

	fmt.Println(count)
}

func B(lines []string) {
	count := 0

	jolts := []int{0}

	for _, l := range lines {
		n, _ := strconv.ParseInt(l, 10, 64)
		jolts = append(jolts, int(n))
	}

	sort.Ints(jolts)

	arr := make([]int, len(jolts), len(jolts))
	arr[len(arr)-1] = 1

	ll := len(arr)

	for i := len(arr) - 2; i >= 0; i-- {
		v := jolts[i]
		sum := 0
		if i+1 < ll && jolts[i+1]-v <= 3 {
			sum += arr[i+1]
		}
		if i+2 < ll && jolts[i+2]-v <= 3 {
			sum += arr[i+2]
		}
		if i+3 < ll && jolts[i+3]-v <= 3 {
			sum += arr[i+3]
		}
		arr[i] = sum
	}

	count = arr[0]

	fmt.Println(count)
}
