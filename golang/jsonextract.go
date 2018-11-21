package main

import (
	"flag"
	"reflect"
	"strings"
)

func main() {
	var keys string
	flag.StringVar(&keys, "keys", "", "json keys to extract")
	flag.Parse()

	if keys == "" {
		return
	}

	keyItems := strings.Split(keys, ",")
	reflect.a
}
