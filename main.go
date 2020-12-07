package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex7"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex7.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex7.A(lines)

		fmt.Print("B: ")
		ex7.B(lines)

		fmt.Println()
	}
}
