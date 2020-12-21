package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex21"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex21.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex21.A(lines)

		fmt.Print("B: ")
		ex21.B(lines)

		fmt.Println()
	}
}
