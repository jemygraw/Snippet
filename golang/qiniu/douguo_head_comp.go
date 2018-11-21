package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

var headTasks chan func()
var initOne sync.Once

func doHead(tasks chan func()) {
	for {
		task := <-tasks
		task()
	}
}

func main() {
	var input string
	var output string
	var worker int
	flag.StringVar(&input, "input", "", "input file path")
	flag.StringVar(&output, "output", "", "output file path")
	flag.IntVar(&worker, "worker", 100, "worker count")
	flag.Parse()

	inputFp, openErr := os.Open(input)
	if openErr != nil {
		fmt.Println("open input file error,", openErr)
		return
	}
	defer inputFp.Close()

	outputFp, openErr := os.Create(output)
	if openErr != nil {
		fmt.Println("open output file error", openErr)
		return
	}
	defer outputFp.Close()

	inputScanner := bufio.NewScanner(inputFp)
	bWriter := bufio.NewWriter(outputFp)

	rwLock := sync.RWMutex{}
	wg := sync.WaitGroup{}
	initOne.Do(func() {
		headTasks = make(chan func(), worker)
		for i := 0; i < worker; i++ {
			go doHead(headTasks)
		}
	})

	for inputScanner.Scan() {
		line := inputScanner.Text()
		items := strings.Split(line, "\t")
		cdnUrl := items[0]
		count := items[1]
		cdnUri, pErr := url.Parse(cdnUrl)
		if pErr != nil {
			continue
		}

		srcUrl := fmt.Sprintf("http://i0.douguo.net%s", cdnUri.Path)
		wg.Add(1)
		headTasks <- func() {
			defer wg.Done()

			var cdnLen int64
			var srcLen int64
			resp, respErr := http.Head(cdnUrl)
			if respErr != nil {
				return
			}

			cdnLen = resp.ContentLength
			resp.Body.Close()

			resp, respErr = http.Head(srcUrl)
			if respErr != nil {
				return
			}

			srcLen = resp.ContentLength
			resp.Body.Close()

			//write
			newLine := fmt.Sprintf("%s\t%s\t%d\t%d\n", cdnUrl, count, cdnLen, srcLen)
			rwLock.Lock()
			bWriter.WriteString(newLine)
			rwLock.Unlock()
		}
	}

	wg.Wait()
	bWriter.Flush()
}
