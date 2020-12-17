package main

import (
	"fmt"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions/ex17"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs(ex17.Path)

	for _, name := range fileNames {
		lines = ioutils.LoadFileLines(name)

		fmt.Println("File: ", name)

		fmt.Print("A: ")
		ex17.A(lines)

		fmt.Print("B: ")
		ex17.B(lines)

		fmt.Println()
	}
}
