package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "/9e55c185d3f2400f0f474d9a29553b98/58c4c8fb/internettv/c1/2016/10/04/1b63c5e12c07b43e21e7718de3727f770.ts"
	items := strings.SplitN(str, "/", 4)
	fmt.Println(items[3])
}
