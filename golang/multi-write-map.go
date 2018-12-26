package main

import (
	"fmt"
	"sync"
	"time"
)

// it is safe to write map like this
func main() {
	fruits := make(map[string][]string)
	fruits["a"] = make([]string, 0)
	fruits["b"] = make([]string, 0)

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fruits["a"] = append(fruits["a"], fmt.Sprintf("apple_%d", i))
			time.Sleep(time.Millisecond * 20)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fruits["b"] = append(fruits["b"], fmt.Sprintf("bear_%d", i))
			time.Sleep(time.Millisecond * 10)
		}
	}()

	wg.Wait()
	fmt.Printf("%v\n", fruits)
}
