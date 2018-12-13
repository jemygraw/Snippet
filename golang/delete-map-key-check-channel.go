package main

import (
	"fmt"
	"time"
)

func main() {
	k := "apple"
	var result bool
	m := make(map[string]chan bool)
	m[k] = make(chan bool)

	go func() {
		defer func() {
			m[k] <- result
			close(m[k])
			delete(m, k)
		}()
	}()

	go func() {
		for {
			select {
			case v := <-m[k]:
				fmt.Println(v)
			default:
				<-time.After(time.Second)
			}
		}
	}()

	result = true

	<-time.After(time.Second * 3)
}
