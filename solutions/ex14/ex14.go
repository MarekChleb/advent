package ex14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const Path = "./solutions/ex14"

type mask []string

func A(lines []string) {
	count := 0

	memory := map[int]int{}
	currentMask := mask(nil)

	for _, l := range lines {
		if isMask(l) {
			currentMask = parseMask(l)
			continue
		}
		k, v := readMem(l)
		memory[k] = currentMask.applyMask(v)
	}

	for _, v := range memory {
		count += v
	}

	fmt.Println(count)
}

func isMask(l string) bool {
	re := regexp.MustCompile("mask = .*")
	return re.MatchString(l)
}

func parseMask(l string) mask {
	re := regexp.MustCompile("mask = (.*)")
	info := re.FindStringSubmatch(l)

	m := mask(strings.Split(info[1], ""))
	return m
}

func readMem(l string) (int, int) {
	re := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)
	info := re.FindStringSubmatch(l)

	k64, _ := strconv.ParseInt(info[1], 10, 64)
	v64, _ := strconv.ParseInt(info[2], 10, 64)

	return int(k64), int(v64)
}

func (m mask) applyMask(v int) int {
	val := 0
	mn := 1
	for i := len(m) - 1; i >= 0; i-- {
		if m[i] == "1" {
			val += mn
		} else if m[i] == "X" {
			if v%2 == 1 {
				val += mn
			}
		}
		mn *= 2
		v /= 2
	}

	return val
}

func B(lines []string) {
	count := 0

	memory := map[int]int{}
	currentMask := mask(nil)

	for _, l := range lines {
		if isMask(l) {
			currentMask = parseMask(l)
			continue
		}
		k, v := readMem(l)
		keys := currentMask.getKeys(k)
		// v2 := currentMask.applyMask(v)
		for _, k2 := range keys {
			memory[k2] = v
		}
	}

	for _, v := range memory {
		count += v
	}

	fmt.Println(count)
}

func (m mask) getKeys(k int) []int {
	val := []int{0}
	app := []int{}
	mn := 1
	for i := len(m) - 1; i >= 0; i-- {
		app = []int{}
		if m[i] == "1" {
			for i, v := range val {
				val[i] = v + mn
			}
		} else if m[i] == "X" {
			for _, v := range val {
				app = append(app, v+mn)
			}
		} else if m[i] == "0" {
			if k%2 == 1 {
				for i, v := range val {
					val[i] = v + mn
				}
			}
		}

		val = append(val, app...)
		mn *= 2
		k /= 2
		// fmt.Println(k, mn)
		// fmt.Println(val)
		// fmt.Println(app)
	}

	return val
}
