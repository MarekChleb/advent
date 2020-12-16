package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex16"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex16.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex16.A(lines)

		fmt.Print("B: ")
		ex16.B(lines)

		fmt.Println()
	}
}
