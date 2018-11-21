package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

//list bucket file
var file string
var dbType string

const (
	DBDataSource = "root:root@(localhost)/redis"
	DBDriver     = "mysql"
	DBMaxConn    = 10
	DBMaxIdle    = 1
)

type FileInfo struct {
	Name     string `xorm:"'name' varchar(200)"`
	Fsize    int64  `xorm:"'fsize'"`
	Hash     string `xorm:"'hash' varchar(28)"`
	PutTime  int64  `xorm:"'put_time'"`
	MimeType string `xorm:"'mime_type' varchar(30)"`
}

const (
	//RedisServer is the local redis server
	RedisServer = "192.168.1.101:6379"
)

func writeRedis(file string) {
	rdsConn, err := redis.Dial("tcp", RedisServer)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rdsConn.Close()

	fp, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	bScanner := bufio.NewScanner(fp)
	nameField := "name"
	fsizeField := "fsize"
	hashField := "hash"
	putTimeField := "putTime"
	mimeTypeField := "mimeType"
	for bScanner.Scan() {
		line := bScanner.Text()
		items := strings.Split(line, "\t")
		name := items[0]
		fsize := items[1]
		hash := items[2]
		putTime := items[3]
		mimeType := items[4]

		rdsKey := fmt.Sprintf("rs1-%s", name)

		_, err := rdsConn.Do("HMSET", rdsKey, nameField, name, fsizeField, fsize, hashField, hash, putTimeField, putTime, mimeTypeField, mimeType)
		if err != nil {
			fmt.Println("Err:", err)
		} else {
			//	fmt.Println(string(reply.(string)))
		}
	}

}

func writeMySQL(file string) {
	dbEngine, err := xorm.NewEngine(DBDriver, DBDataSource)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = dbEngine.Sync2(new(FileInfo))
	if err != nil {
		fmt.Println(err)
		return
	}

	fp, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	bScanner := bufio.NewScanner(fp)

	for bScanner.Scan() {
		line := bScanner.Text()
		items := strings.Split(line, "\t")
		name := items[0]
		fsize, _ := strconv.ParseInt(items[1], 10, 64)
		hash := items[2]
		putTime, _ := strconv.ParseInt(items[3], 10, 64)
		mimeType := items[4]

		fileInfo := FileInfo{
			Name:     name,
			Fsize:    fsize,
			Hash:     hash,
			MimeType: mimeType,
			PutTime:  putTime,
		}

		_, err := dbEngine.Insert(&fileInfo)
		if err != nil {
			fmt.Println("Err:", err)
		}
	}

}

func main() {
	flag.StringVar(&file, "file", "", "file to read")
	flag.StringVar(&dbType, "type", "redis", "redis or mysql")
	flag.Parse()
	start := time.Now()
	fmt.Println("Start:", start)
	if dbType == "mysql" {
		writeMySQL(file)
	} else if dbType == "redis" {
		writeRedis(file)
	} else {
		fmt.Println("Err: wrong type")
	}
	fmt.Println("Duration:", time.Since(start))
}
