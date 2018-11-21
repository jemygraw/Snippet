package main

import (
	"fmt"
)

func main() {

	str := "hello world，金鑫鑫"
	fmt.Println(len([]rune(str)))
	fmt.Println(len([]byte(str)))
}
