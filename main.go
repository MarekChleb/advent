package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex18"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex18.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex18.A(lines)

		fmt.Print("B: ")
		ex18.B(lines)

		fmt.Println()
	}
}
