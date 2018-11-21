package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Println(i, j)
		}
	}
}
