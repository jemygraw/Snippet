package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
223.247.160.139 - 32 [11/Dec/2016:23:43:04 +0800] "GET http://static3.hifun.mobi/video/2016/11/11/dc8d031b7fcdf0801509b9af04d6bb8b.mp4?vframe/jpg/offset/1/w/0/h/0/rotate/auto HTTP/1.1" 200 20469 "-" "Dalvik/2.1.0 (Linux; U; Android 6.0; Letv X500 Build/DBXCNOP5801810092S)"
*/
func main() {
	input := os.Stdin
	bScanner := bufio.NewScanner(input)
	for bScanner.Scan() {
		line := bScanner.Text()
		split(line, " ")
	}
}

func split(line string, delimiter string) {
	next := line

	itemIndex := 0
	for {
		if next == "" {
			//just ignore empty line
			fmt.Println()
			break
		}

		items := strings.SplitN(next, delimiter, 2)
		val := strings.Trim(items[0], "[]\"")
		if itemIndex == 4 {
			val = strings.Replace(val, " ", "\t", -1)
		}

		fmt.Print(val)

		next = strings.TrimSpace(items[1])
		if next != "" {
			fmt.Print("\t")
		}

		if strings.HasPrefix(next, "[") {
			next = next[1:]
			delimiter = "]"
		} else if strings.HasPrefix(next, "\"") {
			next = next[1:]
			delimiter = "\""
		} else {
			delimiter = " "
		}

		itemIndex += 1
	}

}
