package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex6"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex6.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex6.A(lines)

		fmt.Print("B: ")
		ex6.B(lines)

		fmt.Println()
	}
}
