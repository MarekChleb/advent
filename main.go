package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs()

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		solutions.Ex3a(lines)

		fmt.Print("B: ")
		solutions.Ex3b(lines)

		fmt.Println()
	}
}
