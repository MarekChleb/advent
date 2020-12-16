package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex15"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex15.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex15.A(lines)

		fmt.Print("B: ")
		ex15.B(lines)

		fmt.Println()
	}
}
