package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex10"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex10.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex10.A(lines)

		fmt.Print("B: ")
		ex10.B(lines)

		fmt.Println()
	}
}
