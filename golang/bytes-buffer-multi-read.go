package main

import (
	"bytes"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

func main() {
	buffer := bytes.NewBuffer([]byte(nil))
	rand.Seed(time.Now().UnixNano())
	stopSignal := make(chan struct{})
	totalBytes := 0
	var waitSeconds int
	flag.IntVar(&waitSeconds, "wait", 10, "wait seconds")
	flag.Parse()

	go func() {
		<-time.After(time.Second * time.Duration(waitSeconds))
		close(stopSignal)
	}()

	go func() {
		// write to increase buffer size
		for {
			select {
			case <-stopSignal:
				return
			default:
				s := make([]byte, rand.Intn(20))
				rand.Read(s)
				num, _ := buffer.WriteString(base64.StdEncoding.EncodeToString(s))
				totalBytes += num
				<-time.After(time.Millisecond * 10)
			}
		}
	}()

	go func() {
		data := make([]byte, 100)
		var seekStart int64
		fp, _ := os.Create("consumer1.txt")
		for {
			reader := bytes.NewReader(buffer.Bytes())
			reader.Seek(seekStart, io.SeekStart)
			num, _ := reader.Read(data)
			if num > 0 {
				seekStart += int64(num)
				fp.Write(data[:num])
				continue
			}

			select {
			case <-stopSignal:
				return
			default:
				<-time.After(time.Second * 3)
			}
		}
	}()

	go func() {
		data := make([]byte, 100)
		var seekStart int64
		fp, _ := os.Create("consumer2.txt")
		for {
			reader := bytes.NewReader(buffer.Bytes())
			reader.Seek(seekStart, io.SeekStart)
			num, _ := reader.Read(data)
			if num > 0 {
				seekStart += int64(num)
				fp.Write(data[:num])
				continue
			}

			select {
			case <-stopSignal:
				return
			default:
				<-time.After(time.Second * 5)
			}
		}
	}()

	<-stopSignal
	fmt.Println("total bytes", totalBytes, ", wait for ctrl+c!")
	<-time.After(time.Second * 3600)
}
