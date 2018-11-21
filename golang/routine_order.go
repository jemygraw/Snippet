package main

import "fmt"

func main() {
	c := make(chan int, 10000)
	//xx

	go func() {
		fmt.Println("hello world")
		c <- 10
	}()

	go func() {
		fmt.Println("hello ke.com")
		c <- 20
	}()

	fmt.Println(<-c)
	fmt.Println(<-c)
	close(c)
}
