package main

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"
)

func main() {
	input := os.Stdin
	bScanner := bufio.NewScanner(input)
	for bScanner.Scan() {
		line := bScanner.Text()
		uri, pErr := url.Parse(line)
		if pErr != nil {
			continue
		}
		fmt.Println(strings.TrimPrefix(uri.Path, "/"))
	}
}
