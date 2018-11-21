package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var file = "/Users/jemy/Workbase/qshell/bzhan_dianbo/sample.txt"

func main() {
	fp, _ := os.Open(file)
	bReader := bufio.NewReader(fp)
	buffer := make([]byte, 10240)
	for {
		num, err := bReader.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Println(num, err)
	}
}
