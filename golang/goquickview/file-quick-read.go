package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := "./json-encode-decode.go"
	fileContent, readErr := ioutil.ReadFile(filePath)
	if readErr != nil {
		fmt.Println("Err: wow!!!", readErr)
		return
	}
	fmt.Println(string(fileContent))
}
