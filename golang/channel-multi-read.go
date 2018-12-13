package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	buffer := make(chan []byte, 100)
	rand.Seed(time.Now().UnixNano())

	go func() {
		for {
			s := make([]byte, rand.Intn(20))
			rand.Read(s)
			buffer <- []byte(base64.StdEncoding.EncodeToString(s))
			<-time.After(time.Second)
		}
	}()

	go func() {
		for c := range buffer {
			fmt.Println("consumer1:", string(c))
		}
	}()

	go func() {
		for c := range buffer {
			fmt.Println("consumer2:", string(c))
		}
	}()

	<-time.After(time.Second * 3600)
}
