package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	var supportedDomains = []string{
		"http://ottvideoyd.hifuntv.com",
		"http://ottvideoaliyun.hifuntv.com",
		"http://ottvideogs.hifuntv.com",
	}

	var file string
	var worker int
	flag.StringVar(&file, "file", "", "file name")
	flag.IntVar(&worker, "worker", 100, "worker count")
	flag.Parse()

	fp, openErr := os.Open(file)
	if openErr != nil {
		fmt.Println(openErr)
		return
	}
	defer fp.Close()

	statusOkFp, _ := os.Create("tofetch-list.txt")
	statusFailFp, _ := os.Create("nofound-list.txt")
	defer statusFailFp.Close()
	defer statusOkFp.Close()

	wg := sync.WaitGroup{}
	rwLock := sync.RWMutex{}
	counter := 0

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "\t")
		path := items[1]

		if counter%worker == 0 {
			wg.Wait()
		}

		wg.Add(1)

		go func() {
			defer wg.Done()
			var success bool
			for _, domain := range supportedDomains {
				remoteUrl := fmt.Sprintf("%s%s", domain, path)
				resp, respErr := http.Head(remoteUrl)
				if respErr != nil || resp.StatusCode != http.StatusOK {
					success = false
				} else {
					rwLock.Lock()
					items := strings.SplitN(path, "/", 4)
					key := items[3]
					statusOkFp.WriteString(remoteUrl + "\t" + key + "\n")
					rwLock.Unlock()
					success = true
					break
				}
			}

			if !success {
				rwLock.Lock()
				statusFailFp.WriteString(line + "\n")
				rwLock.Unlock()
			}
		}()
	}

	wg.Wait()
}
