package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
)

var gbkEncoder = simplifiedchinese.GBK.NewEncoder()

func utf82GBK(text string) (string, error) {
	return gbkEncoder.String(text)
}

func main() {
	a := "/data/upload/CP2301/qiniustor-sync/七牛.png"
	b, _ := utf82GBK(a)
	fp, openErr := os.Create(b)
	fmt.Println("open error:", openErr)
	fp.Close()

	fi, statErr := os.Stat(b)
	fmt.Println("stat error:", statErr)
	fmt.Println(fi.Name())
}
