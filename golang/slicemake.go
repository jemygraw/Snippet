package main

import (
	"fmt"
	"time"
)

func main() {
	max := 100000000
	start := time.Now()
	s1 := make([]int, 0)
	for i := 0; i < max; i++ {
		s1 = append(s1, i)
	}
	fmt.Println(time.Since(start))
	start = time.Now()
	s2 := make([]int, 0, 1000)
	for i := 0; i < max; i++ {
		s2 = append(s2, i)
	}
	fmt.Println(time.Since(start))
}
