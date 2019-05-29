package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	var file string
	flag.StringVar(&file, "f", "-", "file content to print to console")
	flag.Parse()

	var fp *os.File
	var openErr error
	if file == "-" {
		fp = os.Stdin
	} else {
		fp, openErr = os.Open(file)
		if openErr != nil {
			fmt.Fprint(os.Stderr, openErr.Error())
			return
		}
		defer fp.Close()
	}

	_, cpErr := io.Copy(os.Stdout, fp)
	if cpErr != nil {
		fmt.Fprint(os.Stderr, cpErr.Error())
		return
	}
}
