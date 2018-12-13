package main

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

func main() {
	r, w := io.Pipe()

	go func() {
		defer w.Close()
		i := 0
		for {
			if i > 10 {
				return
			}
			i++
			<-time.After(time.Second * 2)
			fmt.Fprint(w, "i am writing some text for you to eat!\n")
		}
	}()

	go func() {
		for {
			buf := new(bytes.Buffer)
			buf.ReadFrom(r)
			fmt.Print(buf.String())
			<-time.After(time.Second)
		}
	}()
	<-time.After(time.Second * 1000)

}
