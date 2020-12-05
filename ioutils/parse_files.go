package ioutils

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func ReadArgs(path string) (fileNames []string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("[a-zA-Z0-9]+.in$")
	for _, f := range files {
		if re.MatchString(f.Name()) {
			fileNames = append(fileNames, path+"/"+f.Name())
		}
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
