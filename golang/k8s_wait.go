package main

import (
	"context"
	"fmt"
	"time"

	"sync/atomic"

	"k8s.io/apimachinery/pkg/util/wait"
)

func main() {
	g := wait.Group{}
	var counter int32
	ctx, cancelFunc := context.WithCancel(context.Background())
	for i := 0; i < 100; i++ {
		j := i
		g.StartWithContext(ctx, func(ctx context.Context) {
			for {
				//quit if
				if atomic.LoadInt32(&counter) > 1000 {
					cancelFunc() //fire cancel signal
				}
				//otherwise
				select {
				case <-ctx.Done(): //cancel signal received
					return
				default:
					fmt.Println(j, time.Now().String())
					atomic.AddInt32(&counter, 1)
					<-time.After(time.Second)
				}
			}
		})
	}
	g.Wait()
}
