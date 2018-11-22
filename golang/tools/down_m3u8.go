package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	var m3u8Url string
	var worker int
	var destDir string
	flag.StringVar(&m3u8Url, "url", "", "m3u8 url to fetch")
	flag.IntVar(&worker, "worker", 1, "worker count")
	flag.StringVar(&destDir, "dest", "", "save directory")

	flag.Parse()
	if m3u8Url == "" {
		fmt.Println("Error: no m3u8 url")
		return
	}
	if worker <= 0 {
		fmt.Println("Error: invalid worker")
		return
	}
	if destDir == "" {
		fmt.Println("Error: no dest dir")
		return
	}

	resp, respErr := http.Get(m3u8Url)
	if respErr != nil {
		fmt.Println("Error: get m3u8 error,", respErr)
		return
	}
	defer resp.Body.Close()

	m3u8Body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println("Error: read m3u8 error,", readErr)
		return
	}

	// parse m3u8 url
	m3u8URI, _ := url.Parse(m3u8Url)
	m3u8Host := fmt.Sprintf("%s://%s", m3u8URI.Scheme, m3u8URI.Host)

	// save m3u8 content
	m3u8Path := filepath.Join(destDir, filepath.Base(m3u8URI.Path))
	wErr := ioutil.WriteFile(m3u8Path, m3u8Body, 0644)
	if wErr != nil {
		fmt.Println("Error: save m3u8 error", wErr)
		return
	}

	// init downloader
	wg := sync.WaitGroup{}
	once := sync.Once{}
	var tasks chan func()

	once.Do(func() {
		tasks = make(chan func(), worker)
		for i := 0; i < worker; i++ {
			go scheduleTasks(tasks)
		}
	})

	bScanner := bufio.NewScanner(bytes.NewReader(m3u8Body))
	for bScanner.Scan() {
		line := bScanner.Text()
		if strings.HasPrefix(line, "#") {
			// ignore
			continue
		}

		var tsURL string
		var tsPath string
		if strings.HasPrefix(line, "/") {
			//join the path with the m3u8 domain
			tsURL = fmt.Sprintf("%s%s", m3u8Host, line)
			tsPath = filepath.Join(destDir, strings.TrimPrefix(line, "/"))
		} else if strings.HasPrefix(line, "http://") || strings.HasPrefix(line, "https://") {
			//download directly from the ts url
			tsURL = line
			tsURI, _ := url.Parse(line)
			tsPath = filepath.Join(destDir, strings.TrimPrefix(tsURI.Path, "/"))
		} else {
			fmt.Println("Error: invalid ts url", line)
			continue
		}

		// add the download task
		wg.Add(1)
		tasks <- func() {
			defer wg.Done()
			// mkdir if necessary
			localBase := filepath.Dir(tsPath)
			mkErr := os.MkdirAll(localBase, 0755)
			if mkErr != nil {
				fmt.Println("Error: create local dir error,", mkErr)
				return
			}

			localFp, openErr := os.Create(tsPath)
			if openErr != nil {
				fmt.Println("Error: create local file error,", openErr)
				return
			}
			defer localFp.Close()

			resp, respErr := http.Get(tsURL)
			if respErr != nil {
				fmt.Println("Error: download ts error,", respErr)
				return
			}
			defer resp.Body.Close()

			_, cpErr := io.Copy(localFp, resp.Body)
			if cpErr != nil {
				fmt.Println("Error: copy data error,", cpErr)
				return
			}
		}
	}

	// wait for all task ends
	wg.Wait()
}

func scheduleTasks(tasks chan func()) {
	for {
		t := <-tasks
		t()
	}
}
