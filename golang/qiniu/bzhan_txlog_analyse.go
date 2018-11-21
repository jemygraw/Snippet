package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

/*
{
	"httpCode":206,
	"byteSize":8379259,
	"userAgent":"Bilibili Freedoooooom/MarkII",
	"httpMethod":"GET",
	"reqTime":"20180228232147",
	"datasource":"/mnt/storage07/txlog/201802/tx_acgvideo/2018022823-tx.acgvideo.com",
	"province":4,
	"logkit_send_time":"2018-03-24T17:52:31.77248Z",
	"reqIp":"tx.acgvideo.com",
	"referer":"NULL",
	"operator":43,
	"respTime":12036,
	"clientIp":"123.98.28.216",
	"cache":"hit",
	"httpversion":"HTTP/1.1",
	"url":"/12/57/32955712/32955712-1-32.flv?txTime=1519838485&platform=iphone&txSecret=8b3487d244ef11aae6f0ddcecf8800e8&oi=2070027480&rate=394990&hfb=e82b36ca6b2d58b4ab2067d35a6e2480",
	"range":"9550981-"
}
*/

type TencentLog struct {
	HttpCode  int    `json:"httpCode"`
	ByteSize  int    `json:"byteSize"`
	ReqTime   string `json:"reqTime"`
	RespTime  int    `json:"respTime"`
	Cache     string `json:"cache"`
	UserAgent string `json:"userAgent"`
}

func parseDate(dateStr string) (aggDateStr string, err error) {
	srcFmt := "20060102150405"
	srcTime, pErr := time.Parse(srcFmt, dateStr)
	if pErr != nil {
		err = pErr
		return
	}
	dstTime := time.Date(srcTime.Year(), srcTime.Month(), srcTime.Day(), srcTime.Hour(),
		srcTime.Minute()/5*5, 0, 0, time.Local)
	dstTimeFmt := "2006-01-02 15:04:05"
	aggDateStr = dstTime.Format(dstTimeFmt)
	return
}

func main() {
	var file string
	flag.StringVar(&file, "file", "", "bzhan tencent cdn log file")
	flag.Parse()
	if file == "" {
		return
	}

	fp, openErr := os.Open(file)
	if openErr != nil {
		fmt.Println(openErr)
		return
	}
	defer fp.Close()

	//fluxMap := make(map[string]int64)
	uaMap := make(map[string]int64)
	bScanner := bufio.NewScanner(fp)
	for bScanner.Scan() {
		line := bScanner.Text()
		txLog := TencentLog{}
		err := json.Unmarshal([]byte(line), &txLog)
		if err != nil {
			fmt.Println("json err:", line)
			continue
		}

		//fmt.Println(txLog.ReqTime, txLog.Cache, txLog.HttpCode, txLog.ByteSize, txLog.RespTime)
		// aggDate, pErr := parseDate(txLog.ReqTime)
		// if pErr != nil {
		// 	fmt.Println("date err:", line, pErr)
		// 	continue
		// }

		// if _, ok := fluxMap[aggDate]; ok {
		// 	fluxMap[aggDate] += int64(txLog.ByteSize)
		// } else {
		// 	fluxMap[aggDate] = int64(txLog.ByteSize)
		// }

		if _, ok := uaMap[txLog.UserAgent]; ok {
			uaMap[txLog.UserAgent] += 1
		} else {
			uaMap[txLog.UserAgent] = 1
		}
	}

	// for k, v := range fluxMap {
	// 	fmt.Println(k, v)
	// }

	for k, v := range uaMap {
		fmt.Println(k, v)
	}

}
