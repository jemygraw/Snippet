package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "http://"
	_, err := url.Parse(str)
	if err != nil {
		fmt.Println(err)
	}
}
