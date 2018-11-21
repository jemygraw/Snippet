package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	var file string
	flag.StringVar(&file, "file", "", "url list file")
	flag.Parse()
	listFp, openErr := os.Open(file)
	if openErr != nil {
		fmt.Println(openErr)
		return
	}

	defer listFp.Close()

	bScanner := bufio.NewScanner(listFp)
	for bScanner.Scan() {
		line := bScanner.Text()

		resp, respErr := http.Head(line)
		if respErr != nil {
			fmt.Println(line + "\t0")
			continue
		}
		fmt.Printf("%s\t%d\n", line, resp.ContentLength)
		resp.Body.Close()
	}
}
