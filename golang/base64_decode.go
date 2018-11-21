package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "4oCN4oCN4oCN"
	b, _ := base64.URLEncoding.DecodeString(str)
	fmt.Println(b)
	fmt.Println(string(b))
}
