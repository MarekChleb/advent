package ex16

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const Path = "./solutions/ex16"

type rangeValidator func(v int) bool
type ticket []int

func A(lines []string) {
	count := 0

	rangeNames := []string{}
	rangeValidators := []rangeValidator{}
	isMyTicket := false
	myTicket := ticket{}

	isNearbyTickets := false
	tickets := []ticket{}

	for _, l := range lines {
		if isRange(l) {
			name, validator := getRange(l)
			rangeNames = append(rangeNames, name)
			rangeValidators = append(rangeValidators, validator)
			continue
		}
		if l == "" {
			continue
		}
		if l == "your ticket:" {
			isMyTicket = true
			continue
		}
		if isMyTicket {
			myTicket = parseTicket(l)
			isMyTicket = false
			continue
		}
		if l == "nearby tickets:" {
			isNearbyTickets = true
			continue
		}
		if isNearbyTickets {
			t := parseTicket(l)
			tickets = append(tickets, t)
		}
	}

	for _, t := range tickets {
		count += t.badValuesSum(rangeValidators)
	}

	fmt.Println(myTicket)
	// fmt.Println(tickets)

	fmt.Println(count)
}

func isRange(s string) bool {
	reRange := regexp.MustCompile(`(.*): (.*)`)

	return reRange.MatchString(s)
}

func getRange(s string) (string, rangeValidator) {
	reRange := regexp.MustCompile(`(.*): (.*)$`)

	info := reRange.FindStringSubmatch(s)
	name := info[1]

	reRangesValues := regexp.MustCompile(`(\d+)-(\d+) or (\d+)-(\d+)`)
	infoValues := reRangesValues.FindStringSubmatch(info[2])

	vals := []int{}
	for _, vs := range infoValues[1:] {
		v, _ := strconv.ParseInt(vs, 10, 64)
		vals = append(vals, int(v))
	}

	validator := createRangeValidator(vals[0], vals[1], vals[2], vals[3])

	return name, validator
}

func createRangeValidator(a, b, c, d int) rangeValidator {
	return func(v int) bool {
		return (v >= a && v <= b) || (v >= c && v <= d)
	}
}

func parseTicket(s string) ticket {
	sVals := strings.Split(s, ",")

	t := ticket{}

	for _, vs := range sVals {
		v, _ := strconv.ParseInt(vs, 10, 64)
		t = append(t, int(v))
	}

	return t
}

func (t ticket) badValuesSum(validators []rangeValidator) int {
	sum := 0

	for _, v := range t {
		isValid := false
		for _, validator := range validators {
			if validator(v) {
				isValid = true
			}
		}

		if !isValid {
			sum += v
		}
	}

	return sum
}

func B(lines []string) {
	count := 0

	rangeNames := []string{}
	rangeValidators := []rangeValidator{}
	isMyTicket := false
	myTicket := ticket{}

	isNearbyTickets := false
	tickets := []ticket{}

	for _, l := range lines {
		if isRange(l) {
			name, validator := getRange(l)
			rangeNames = append(rangeNames, name)
			rangeValidators = append(rangeValidators, validator)
			continue
		}
		if l == "" {
			continue
		}
		if l == "your ticket:" {
			isMyTicket = true
			continue
		}
		if isMyTicket {
			myTicket = parseTicket(l)
			isMyTicket = false
			continue
		}
		if l == "nearby tickets:" {
			isNearbyTickets = true
			continue
		}
		if isNearbyTickets {
			t := parseTicket(l)
			tickets = append(tickets, t)
		}
	}

	validTickets := []ticket{}

	// if myTicket.badValuesSum(rangeValidators) == 0 {
	// 	validTickets = append(validTickets, myTicket)
	// }

	for _, t := range tickets {
		if t.badValuesSum(rangeValidators) == 0 {
			validTickets = append(validTickets, t)
		}
	}

	mapping := getMapping(validTickets, rangeValidators)

	fmt.Println(mapping)

	count = 1
	for i, name := range rangeNames {
		nameSplit := strings.Split(name, " ")
		if nameSplit[0] == "departure" {
			count *= myTicket[mapping[i]]
		}
	}

	// fmt.Println(myTicket)
	// fmt.Println(tickets)

	fmt.Println(count)
}

func getMapping(validTickets []ticket, validators []rangeValidator) []int {
	mapping := make([]int, len(validators))

	found := map[int]int{}

	ticketLen := len(validTickets[0])

	possibleFields := map[int][]int{}

	for vi, validator := range validators {
		possibleFields[vi] = []int{}
		for i := 0; i < ticketLen; i++ {
			if _, ok := found[i]; ok {
				continue
			}
			isValid := true
			for _, t := range validTickets {
				if !validator(t[i]) {
					isValid = false
					break
				}
			}
			if isValid {
				possibleFields[vi] = append(possibleFields[vi], i)
			}
		}
	}

	fmt.Println(possibleFields)

	sortedValidators := []int{}
	for i := range validators {
		sortedValidators = append(sortedValidators, i)
	}

	sort.Slice(sortedValidators, func(i, j int) bool {
		return len(possibleFields[sortedValidators[i]]) < len(possibleFields[sortedValidators[j]])
	})

	for _, i := range sortedValidators {
		poss := -1
		for _, v := range possibleFields[i] {
			if _, ok := found[v]; !ok {
				poss = v
				break
			}
		}
		found[poss] = i
	}

	for i := range validators {
		mapping[found[i]] = i
	}

	return mapping
}
