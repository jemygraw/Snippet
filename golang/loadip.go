package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

const MaxIPInt int = 255

func doGet(tasks chan func()) {
	for {
		task := <-tasks
		task()
	}
}

func main() {
	var ipCmd string
	var worker int
	var output string

	var taskChan chan func()

	flag.IntVar(&worker, "worker", 100, "worker count")
	flag.StringVar(&output, "output", "ips.txt", "output file")
	flag.StringVar(&ipCmd, "cmd", "", "ip command path")

	flag.Parse()
	if ipCmd == "" {
		fmt.Println("please specifiy the ip command")
		return
	}

	outputFp, openErr := os.OpenFile(output, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if openErr != nil {
		fmt.Println("open output file error", openErr)
		return
	}
	defer outputFp.Close()

	bWriter := bufio.NewWriter(outputFp)
	defer bWriter.Flush()

	wg := sync.WaitGroup{}
	rwLock := sync.RWMutex{}

	var initOne sync.Once
	initOne.Do(func() {
		taskChan = make(chan func(), worker)
		for i := 0; i < worker; i++ {
			go doGet(taskChan)
		}
	})

	for k1 := 1; k1 <= MaxIPInt; k1++ {
		for k2 := 1; k2 <= MaxIPInt; k2++ {
			for k3 := 1; k3 <= MaxIPInt; k3++ {
				for k4 := 1; k4 <= MaxIPInt; k4++ {
					newIP := fmt.Sprintf("%d.%d.%d.%d", k1, k2, k3, k4)

					wg.Add(1)
					taskChan <- func() {
						defer wg.Done()
						execCmd := exec.Command(ipCmd, newIP)
						execOuput, execErr := execCmd.Output()
						if execErr != nil {
							fmt.Println("get ip", newIP, "error", execErr)
							return
						}

						line := fmt.Sprintf("%s %s", newIP, string(execOuput))

						rwLock.Lock()
						bWriter.WriteString(line)
						rwLock.Unlock()
					}
				}
			}
		}
	}

	wg.Wait()
}
