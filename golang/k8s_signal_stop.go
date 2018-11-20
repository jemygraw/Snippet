package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

var onlyOneSignalHandler = make(chan struct{})

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandler() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}

func main() {
	stop := SetupSignalHandler()
	//	SetupSignalHandler()
	for {
		select {
		case <-stop:
			fmt.Println("gracefully exit...")
			<-time.After(time.Second * 5)
			fmt.Println("i am gone...")
			os.Exit(0)
		default:
			fmt.Println("i am running happily...")
			<-time.After(time.Second)
		}
	}
}
