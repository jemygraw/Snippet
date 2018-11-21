package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func getint() int {
	return 100
}

func main() {
	var times int
	flag.IntVar(&times, "times", 1, "times to call syscall")
	flag.Parse()
	start := time.Now()
	for i := 0; i < times; i++ {
		os.Getpid()
	}

	fmt.Println(time.Since(start))

	start = time.Now()
	for i := 0; i < times; i++ {
		getint()
	}
	fmt.Println(time.Since(start))
}
