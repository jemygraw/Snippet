package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	var msg = `Jul 26 17:18:43 tiger com.apple.xpc.launchd[1] (com.alipay.DispatcherService): Service only ran for 2 seconds. Pushing respawn out by 8 seconds.`

	conn, err := net.Dial("udp", "127.0.0.1:5140")
	fmt.Println(err)
	for {
		_, err = conn.Write([]byte(msg))
		fmt.Println(err)
		<-time.After(time.Second * 1)
	}
}
