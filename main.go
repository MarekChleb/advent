package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mw332637/advent/ioutils"
	"github.com/mw332637/advent/solutions"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs()

	for _, name := range fileNames {
		fmt.Println("File: ", name)
		lines = []string{}
		file, _ := os.Open(name)
		r := bufio.NewScanner(file)
		for r.Scan() {
			lines = append(lines, r.Text())
		}
		file.Close()

		fmt.Print("A: ")
		solutions.Ex3a(lines)

		fmt.Print("B: ")
		solutions.Ex3b(lines)

		fmt.Println()
	}
}
