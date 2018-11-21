package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("pid:", os.Getpid())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, os.Interrupt, os.Kill, syscall.SIGQUIT)
	s := <-c
	fmt.Println("cancel", s)

}
