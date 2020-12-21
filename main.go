package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex19"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex19.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex19.A(lines)

		fmt.Print("B: ")
		ex19.B(lines)

		fmt.Println()
	}
}
