package main

import (
	"fmt"
	"sync"
	"time"
)

/*
对于并发执行的程序，fmt.Println 输出的时候最好带上 time.Unix().UnixNano() 时间戳，
这样对输出的结果进行 sort $1 排序，然后才方便看出指令的执行顺序。

$ go run rwdemo.go |sort $1
*/
func main() {
	wg := sync.WaitGroup{}
	rwLock := sync.RWMutex{}
	//可以通过调节这两个limit参数来体会锁定的机制
	limitRwlock := 2
	limitLock := 10

	for i := 0; i < limitLock; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(time.Now().UnixNano(), "try to rlock", i, "@")
			rwLock.RLock()
			fmt.Println(time.Now().UnixNano(), "rlock", i, "!")
			time.Sleep(time.Second * 2)
			rwLock.RUnlock()
			fmt.Println(time.Now().UnixNano(), "unrlock", i, "~")
		}(i)
	}

	for i := 0; i < limitRwlock; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(time.Now().UnixNano(), "try to rwlock", i, "@")
			rwLock.Lock()
			fmt.Println(time.Now().UnixNano(), "rwlock", i, "!")
			time.Sleep(time.Second * 2)
			rwLock.Unlock()
			fmt.Println(time.Now().UnixNano(), "rwunlock", i, "~")
		}(i)
	}

	wg.Wait()
}
