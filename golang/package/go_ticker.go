package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Tick(time.Millisecond * 200)
	for e := range t {
		fmt.Println(e)
	}
}
