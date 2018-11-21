package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"golang.org/x/net/websocket"
)

func main() {
	var origin string
	var url string
	flag.StringVar(&origin, "origin", "", "websocket origin")
	flag.StringVar(&url, "url", "", "websocket remote url")
	flag.Parse()

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		panic(err)
		return
	}

	buffer := make([]byte, 100)
	bScanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for bScanner.Scan() {
		line := bScanner.Text()
		if line == "" {
			fmt.Print("> ")
			continue
		}
		ws.Write([]byte(line + "\r\n"))

		totalReadBytes := 0
		remainedBytes := len(buffer)
		for remainedBytes > 0 {
			num, err := ws.Read(buffer)
			if err != nil {
				if err == io.EOF {
					fmt.Print(string(buffer[:num]))
					break
				}
				ws.Close()
				return
			}
			totalReadBytes += num
			remainedBytes = ws.Len() - totalReadBytes

			fmt.Print(string(buffer[:num]))
		}

		// for next command
		fmt.Println()
		fmt.Print("> ")
	}
}
