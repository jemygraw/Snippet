package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 100)
	for i := 0; i < 1000; i++ {
		fmt.Println(i)
		c <- i
	}

	<-time.After(time.Second * 10)
}
