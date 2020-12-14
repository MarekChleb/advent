package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex14"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex14.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex14.A(lines)

		fmt.Print("B: ")
		ex14.B(lines)

		fmt.Println()
	}
}
