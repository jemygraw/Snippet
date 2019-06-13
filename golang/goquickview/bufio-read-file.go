package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "./json-encode-decode.go"
	fp, openErr := os.Open(filePath)
	if openErr != nil {
		fmt.Println("Err: wow!!!", openErr)
		return
	}
	defer fp.Close()

	// use buffer reader to read each line
	bScanner := bufio.NewScanner(fp)
	for bScanner.Scan() {
		line := bScanner.Text()
		fmt.Println(line)
	}
}
