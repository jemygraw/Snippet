package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(len(strings.SplitN("hello", " ", 2)))
}
