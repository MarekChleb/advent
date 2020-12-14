package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex13"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex13.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex13.A(lines)

		fmt.Print("B: ")
		ex13.B(lines)

		fmt.Println()
	}
}
