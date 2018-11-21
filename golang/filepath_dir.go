package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	s := "/a/b/c/d.m3u8"
	fmt.Println(filepath.Dir(s))
}
