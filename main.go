package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex9"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex9.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex9.A(lines)

		fmt.Print("B: ")
		ex9.B(lines)

		fmt.Println()
	}
}
