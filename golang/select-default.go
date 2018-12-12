package main

import (
	"fmt"
	"time"
)

func main() {
	s := make(chan bool)

	go func() {
		for {
			select {
			case <-s:
				fmt.Println("stop now")
				return
			default:
				func() {
					<-time.After(time.Second * 10)
					fmt.Println("wait for 10 seconds")
				}()
			}
		}
	}()
	<-time.After(time.Second * 30)
	s <- true
	<-time.After(time.Second * 10)
}
