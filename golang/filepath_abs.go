package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	dir := "."
	fname := "hello.go"
	newDir := filepath.Join(dir, fname)
	fmt.Println(newDir)
	newDirAbs, _ := filepath.Abs(newDir)
	fmt.Println(newDirAbs)
}
