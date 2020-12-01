package ioutils

import "os"

func ReadArgs() (fileNames []string) {
	for i := 1; i < len(os.Args); i++ {
		fileNames = append(fileNames, os.Args[i])
	}

	return
}
