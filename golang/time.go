package main

import (
	"fmt"
	"time"
)

func main() {
	var t int64 = 1480780729831616 * 1000
	fmt.Println(time.Unix(0, t).Format("16/Jul/2016:20:00:00 +0800"))

}
