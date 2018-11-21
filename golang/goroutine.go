package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("i am an unhappy go routine...")
	}()

	fmt.Println("app exits")
	//<-time.After(time.Second)
}
