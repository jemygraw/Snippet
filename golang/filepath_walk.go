package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	srcDir := "../"
	fmt.Println(filepath.Abs(srcDir))
	filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) (runErr error) {
		fmt.Println(filepath.Abs(path))
		return
	})
}
