package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex5"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex5.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex5.A(lines)

		fmt.Print("B: ")
		ex5.B(lines)

		fmt.Println()
	}
}
