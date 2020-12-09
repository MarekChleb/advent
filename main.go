package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex8"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex8.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex8.A(lines)

		fmt.Print("B: ")
		ex8.B(lines)

		fmt.Println()
	}
}
