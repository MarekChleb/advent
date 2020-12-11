package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex11"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex11.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex11.A(lines)

		fmt.Print("B: ")
		ex11.B(lines)

		fmt.Println()
	}
}
