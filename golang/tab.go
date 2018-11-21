package main

import (
	"fmt"
	"strings"
)

func main() {
	sep := "\\t"
	fmt.Println("sep", sep)

	//fix
	if sep == "\\t" {
		sep = "\t"
	}

	//split
	str := "hello	world	qiniu	cloud	storage"
	items := strings.Split(str, sep)
	for i, v := range items {
		fmt.Println(i, v)
	}
}
