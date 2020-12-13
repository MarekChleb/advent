package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex12"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex12.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex12.A(lines)

		fmt.Print("B: ")
		ex12.B(lines)

		fmt.Println()
	}
}
