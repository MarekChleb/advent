package ex8

import (
	"fmt"
	"strconv"
	"strings"
)

const Path = "./solutions/ex8"

type Operation struct {
	Type     string
	Argument int
	Visited  bool
}

func A(lines []string) {
	operations := []Operation{}

	for _, l := range lines {
		operations = append(operations, parseOperation(l))
	}

	count := Perform(operations)

	fmt.Println(count)
}

func B(lines []string) {
	count := 0
	operations := []Operation{}
	for _, l := range lines {
		operations = append(operations, parseOperation(l))
	}

	count, _ = PerformWithError(&operations, 0, 0, false)

	fmt.Println(count)
}

func parseOperation(l string) Operation {
	lSplitted := strings.Split(l, " ")
	arg, _ := strconv.ParseInt(lSplitted[1], 10, 64)

	return Operation{Type: lSplitted[0], Argument: int(arg)}
}

func Perform(operations []Operation) int {
	index := 0
	acc := 0
	for {
		op := operations[index]
		oldIndex := index
		if op.Visited {
			return acc
		}
		switch op.Type {
		case "nop":
			index++
			break
		case "acc":
			index++
			acc += op.Argument
			break
		case "jmp":
			index += op.Argument
			break
		}
		operations[oldIndex] = Operation{
			Type:     op.Type,
			Argument: op.Argument,
			Visited:  true,
		}
	}
}

func PerformWithError(operations *[]Operation, currentAcc, currentIndex int, wasError bool) (int, bool) {
	if currentIndex >= len(*operations) {
		// fmt.Printf("%d %v\n", currentIndex, wasError)
		return currentAcc, wasError
	}
	if currentIndex < 0 {
		return currentAcc, false
	}

	op := (*operations)[currentIndex]
	if op.Visited {
		return currentAcc, false
	}

	var ok bool
	var acc int

	defer func() {
		if ok {
			// fmt.Printf("%d ", currentIndex)
		}
	}()

	switch op.Type {
	case "nop":
		if !wasError {
			(*operations)[currentIndex] = Operation{"jmp", op.Argument, true}
			acc, ok = PerformWithError(operations, currentAcc, currentIndex+op.Argument, true)
			if ok {
				return acc, ok
			}
		}
		(*operations)[currentIndex] = Operation{"nop", op.Argument, true}
		acc, ok = PerformWithError(operations, currentAcc, currentIndex+1, wasError)
		if ok {
			return acc, ok
		}
		break
	case "jmp":
		if !wasError {
			(*operations)[currentIndex] = Operation{"nop", op.Argument, true}
			acc, ok = PerformWithError(operations, currentAcc, currentIndex+1, true)
			if ok {
				return acc, ok
			}
		}
		(*operations)[currentIndex] = Operation{"jmp", op.Argument, true}
		acc, ok = PerformWithError(operations, currentAcc, currentIndex+op.Argument, wasError)
		if ok {
			return acc, ok
		}
		break
	case "acc":
		(*operations)[currentIndex] = Operation{"acc", op.Argument, true}
		acc, ok = PerformWithError(operations, currentAcc+op.Argument, currentIndex+1, wasError)
		if ok {
			return acc, ok
		}
		break
	}

	(*operations)[currentIndex] = Operation{op.Type, op.Argument, false}
	return acc, false
}
