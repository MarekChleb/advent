package ioutils

import (
	"bufio"
	"os"
)

func ReadArgs() (fileNames []string) {
	for i := 1; i < len(os.Args); i++ {
		fileNames = append(fileNames, os.Args[i])
	}

	return
}

func LoadFileLines(filename string) []string {
	lines := []string{}
	file, _ := os.Open(filename)
	r := bufio.NewScanner(file)
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	file.Close()
	return lines
}
