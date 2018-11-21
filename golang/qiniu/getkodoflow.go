package main

import (
	"fmt"
	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
	"time"
)

const (
	BUCKET_HUADONG = "z0"
	BUCKET_HUABEI  = "z1"
	BUCKET_HUANAN  = "z2"
	BUCKET_BEIMEI  = "na0"
)

const (
	FLOW_API = "http://api.qiniu.com/v6/blob_io"
)

const (
	AccessKey = ""
	SecretKey = ""
)

type FlowInfo struct {
	Time   time.Time `json:"time"`
	Values Flow      `json:"values"`
}

type Flow struct {
	Flow int64 `json:"flow"`
}

/*
beginDate - like 20161216
beginTime - like 000000
endDate - like 20161216
endTime - like 235959

@optional params
interval - 5min or day
bucket - bucket name
bucketZone - bucket zone
srcDomain - src supplier domain
srcSupplier - src supplier name, can be [origin, inner, ws, tx, tx_ov, kw, letv, tc, dl, bs, unknown]
*/
func getFlowInfo(beginDate, beginTime, endDate, endTime, interval, bucket, bucketZone, srcDomain, srcSupplier string) (flowInfos []FlowInfo, err error) {
	c := kodo.New(0, nil)
	ctx := context.TODO()

	if interval == "" {
		interval = "day"
	}

	params := map[string][]string{
		"begin":  []string{fmt.Sprintf("%s/%s", beginDate, beginTime)},
		"select": []string{"flow"},
		"end":    []string{fmt.Sprintf("%s/%s", endDate, endTime)},
		"g":      []string{interval},
	}

	//bucket info
	if bucket != "" {
		params["$bucket"] = []string{bucket}
	}

	if bucketZone != "" {
		params["$zone"] = []string{bucketZone}
	}

	//supplier
	if srcSupplier != "" {
		params["$src"] = []string{srcSupplier}
	}

	if srcDomain != "" {
		params["$domain"] = []string{srcDomain}
	}

	callUrl := fmt.Sprintf("%s", FLOW_API)
	callErr := c.CallWithForm(ctx, &flowInfos, "GET", callUrl, params)
	if callErr != nil {
		err = fmt.Errorf("get space info error, %s", callErr)
		return
	}

	return
}

func main() {

	conf.ACCESS_KEY = AccessKey
	conf.SECRET_KEY = SecretKey

	beginDate := "20161216"
	beginTime := "00:00"
	endDate := "20161218"
	endTime := "00:00"

	interval := "day"
	//set as you need, if none, let it empty
	bucket := ""
	bucketZone := ""
	srcDomain := ""
	srcSupplier := ""

	flowInfos, gErr := getFlowInfo(beginDate, beginTime, endDate, endTime, interval, bucket, bucketZone, srcDomain, srcSupplier)
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	if len(flowInfos) == 0 {
		fmt.Println("get flow info error, no data found")
		return
	}

	for _, info := range flowInfos {
		fmt.Println(info.Time.Unix(), info.Values.Flow)
	}
}
