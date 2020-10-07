package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var files []*os.File
	if len(os.Args) == 1 {
		files = append(files, os.Stdin)
	} else {
		for _, fileName := range os.Args[1:] {
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to open file: %v\n", err)
				os.Exit(1)
			}
			defer file.Close()
			files = append(files, file)
		}
	}

	for _, file := range files {
		_, err := io.Copy(os.Stdout, file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write to stdout: %v\n", err)
			os.Exit(1)
		}
	}
}
