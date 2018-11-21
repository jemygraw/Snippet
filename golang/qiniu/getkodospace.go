package main

import (
	"fmt"
	"golang.org/x/net/context"
	"qiniupkg.com/api.v7/conf"
	"qiniupkg.com/api.v7/kodo"
)

const (
	BUCKET_HUADONG = "z0"
	BUCKET_HUABEI  = "z1"
	BUCKET_HUANAN  = "z2"
	BUCKET_BEIMEI  = "na0"
)

const (
	SPACE_API = "http://api.qiniu.com/v6/space"
)

const (
	AccessKey = ""
	SecretKey = ""
)

/*
beginDate - like 20161216
beginTime - like 000000
endDate - like 20161216
endTime - like 235959
interval - 5min or day
bucket - optional, bucket name
bucektZone - bucket zone
*/
func getSpaceInfo(beginDate, beginTime, endDate, endTime, interval, bucket, bucketZone string) (spaceInfo map[string][]int64, err error) {
	c := kodo.New(0, nil)
	ctx := context.TODO()
	params := map[string][]string{
		"begin": []string{fmt.Sprintf("%s%s", beginDate, beginTime)},
		"end":   []string{fmt.Sprintf("%s%s", endDate, endTime)},
		"g":     []string{interval},
	}

	if bucket != "" {
		params["bucket"] = []string{bucket}
	}

	callUrl := fmt.Sprintf("%s/%s", SPACE_API, bucketZone)
	callErr := c.CallWithForm(ctx, &spaceInfo, "GET", callUrl, params)
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
	beginTime := "000000"
	endDate := "20161218"
	endTime := "000000"
	interval := "day"
	bucket := "" //set as you need, if none, let it empty
	bucketZone := BUCKET_HUANAN
	spaceInfo, gErr := getSpaceInfo(beginDate, beginTime, endDate, endTime, interval, bucket, bucketZone)
	if gErr != nil {
		fmt.Println(gErr)
		return
	}

	times := spaceInfo["times"]
	datas := spaceInfo["datas"]

	if len(times) == 0 {
		fmt.Println("get space info error, no data found")
		return
	}

	for i := 0; i < len(times); i++ {
		fmt.Printf("%d\t%d\n", times[i], datas[i])
	}
}
