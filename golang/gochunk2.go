package main

import (
	"fmt"
	"regexp"
)

func main() {
	path := "/upload/dish/$1/$2/$3/$5"
	pattern := `\$\d+`
	rgxp := regexp.MustCompile(pattern)
	fmt.Println(rgxp.FindAllString(path, -1))
}
