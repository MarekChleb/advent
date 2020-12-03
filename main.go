package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mw332637/advent/ioutils"
)

var lines []string

func main() {
	fileNames := ioutils.ReadArgs()

	for _, name := range fileNames {
		fmt.Println("\nFile: ", name)
		lines = []string{}
		file, _ := os.Open(name)
		r := bufio.NewScanner(file)
		for r.Scan() {
			lines = append(lines, r.Text())
		}
		file.Close()

		ex3a()
		ex3b()
	}
}
