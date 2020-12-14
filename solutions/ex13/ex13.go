package ex13

import (
	"fmt"
	"strconv"
	"strings"
)

const Path = "./solutions/ex13"

func A(lines []string) {
	count := 0

	timeStr := lines[0]

	time, _ := strconv.ParseInt(timeStr, 10, 64)
	min := int64(1000000000000)

	mods := strings.Split(lines[1], ",")

	for _, modStr := range mods {
		if modStr == "x" {
			continue
		}
		mod, _ := strconv.ParseInt(modStr, 10, 64)
		modV := time % mod

		if modV == 0 {
			count = 0
			break
		}

		plus := mod - modV
		if min > time+plus {
			min = time + plus
			count = int(plus * mod)
		}
	}

	fmt.Println(count)
}

func B(lines []string) {
	count := int64(0)

	mods := strings.Split(lines[1], ",")

	str := []string{}
	cmod := int64(1)

	for i, modStr := range mods {
		if modStr != "x" {
			mod, _ := strconv.ParseInt(modStr, 10, 64)
			sum := int64(i)

			num := int64(i) % mod
			for {
				if sum%mod == num {
					count = sum
					cmod *= mod
					fmt.Printf("%d mod %d = 0\n", sum, cmod)
					break
				}
				sum += cmod
			}

			// str = append(str, fmt.Sprintf("((x+%d) mod %d=0)", i, mod))
			str = append(str, fmt.Sprintf("(x mod %d=%d)", mod, num))
		}
	}

	fmt.Println(count)
	fmt.Println(strings.Join(str, " & "))
}
