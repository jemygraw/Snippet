package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	s := "/Users/jemy/Documents/qiniu.png"
	fmt.Println(filepath.Base(s))
	fmt.Println(filepath.Dir(s))
}
