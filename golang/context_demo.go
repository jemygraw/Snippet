package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("work%d", i)
		wg.Add(1)
		go func() {
			defer wg.Done()
			work(ctx, name)
		}()
	}
	<-time.After(time.Second * 10)
	cancelFunc()

	//wait for all the work done
	wg.Wait()
}

func work(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			//quit the work when done() fired
			fmt.Println("quit", name, ctx.Err())
			return
		default:
			<-time.After(time.Second * 1)
		}
		fmt.Println(name, time.Now().String())
	}
}
