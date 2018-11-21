package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"k8s.io/apimachinery/pkg/util/wait"
)

var engine *xorm.Engine

func ping(engine *xorm.Engine) {
	wait.Forever(func() {
		err := engine.Ping()
		if err != nil {
			fmt.Println(err)
		}
	}, time.Second*10)

}

func main() {
	var ch chan int
	var err error
	ch = make(chan int)
	engine, err = xorm.NewEngine("mysql", "root:root@(localhost:3306)/atlas?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	engine.ShowSQL(true)
	engine.ShowExecTime(true)
	go ping(engine)
	<-ch
}
