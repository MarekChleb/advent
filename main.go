package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex4"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex4.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex4.A(lines)

		fmt.Print("B: ")
		ex4.B(lines)

		fmt.Println()
	}
}
