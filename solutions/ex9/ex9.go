package ex9

import (
	"fmt"
	"strconv"
)

const Path = "./solutions/ex9"

func A(lines []string) {
	count := 0

	plen := 25
	pam := []int{}
	for i := 0; i < plen; i++ {
		ns := lines[i]
		num, _ := strconv.ParseInt(ns, 10, 64)
		pam = append(pam, int(num))
	}

	for i := plen; i < len(lines); i++ {
		ns := lines[i]
		num, _ := strconv.ParseInt(ns, 10, 64)
		ok := check(pam, int(num))
		pam = append(pam[1:], int(num))
		if !ok {
			count = int(num)
			break
		}
	}

	fmt.Println(count)
}

func B(lines []string) {
	count := 0

	find := 88311122
	pam := []int{}
	for i := 0; i < len(lines); i++ {
		ns := lines[i]
		num, _ := strconv.ParseInt(ns, 10, 64)

		if int(num) == find {
			break
		}
		pam = append(pam, int(num))
	}

	count = getSum(pam, find)

	fmt.Println(count)
}

func check(pam []int, num int) bool {
	for i, v := range pam {
		for j, v2 := range pam {
			if i < j {
				if v+v2 == num {
					return true
				}
			}
		}
	}
	return false
}

func getSum(pam []int, find int) int {
	sum := 0

	for l := 2; l < len(pam); l++ {
		for e := l - 1; e < len(pam); e++ {
			for i := e - l + 1; i < e+1; i++ {
				sum += pam[i]
			}
			if sum == find {
				return sum2(pam, e-l+1, e)
			}
			sum = 0
		}
	}

	return 0
}

func sum2(pam []int, l, r int) int {
	min := 1123120120
	max := -123
	for i := l; i <= r; i++ {
		if min > pam[i] {
			min = pam[i]
		}
		if max < pam[i] {
			max = pam[i]
		}
	}

	return min + max
}
