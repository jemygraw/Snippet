package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/qiniu/api.v6/auth/digest"
	"github.com/qiniu/rpc"
	"net/http"
	"os"
	"strings"
)

const (
	PfopUrl   = "http://api.qiniu.com/pfop/"
	PrefopUrl = "http://api.qiniu.com/status/get/prefop?id=%s"
	NotifyUrl = "xxx"
	AccessKey = "xxx"
	SecretKey = "xxx"
)

type PrefopRet struct {
	Id          string `json:"id"`
	Code        int    `json:"code"`
	Desc        string `json:"desc"`
	InputBucket string `json:"inputBucket,omitempty"`
	InputKey    string `json:"inputKey,omitempty"`
	Pipeline    string `json:"pipeline,omitempty"`
	Reqid       string `json:"reqid,omitempty"`
	Items       []PrefopResult
}
type PrefopResult struct {
	Cmd   string   `json:"cmd"`
	Code  int      `json:"code"`
	Desc  string   `json:"desc"`
	Error string   `json:"error,omitempty"`
	Hash  string   `json:"hash,omitempty"`
	Key   string   `json:"key,omitempty"`
	Keys  []string `json:"keys,omitempty"`
}

type PfopResult struct {
	PersistentId string `json:"persistentId,omitempty"`
}

func main() {
	var file string
	flag.StringVar(&file, "f", "", "persistent id list file")
	flag.Parse()

	fp, openErr := os.Open(file)
	if openErr != nil {
		fmt.Println(openErr)
		return
	}
	defer fp.Close()

	mac := digest.Mac{
		AccessKey,
		[]byte(SecretKey),
	}

	t := digest.NewTransport(&mac, nil)
	client := &http.Client{Transport: t}
	rpcClient := rpc.Client{client}

	bScanner := bufio.NewScanner(fp)
	bScanner.Split(bufio.ScanLines)
	for bScanner.Scan() {
		pid := bScanner.Text()
		func() {
			prefopUrl := fmt.Sprintf(PrefopUrl, pid)
			resp, respErr := http.Get(prefopUrl)
			if respErr != nil {
				fmt.Println(respErr, pid)
				return
			}
			defer resp.Body.Close()

			prefopRet := PrefopRet{}
			decoder := json.NewDecoder(resp.Body)
			mErr := decoder.Decode(&prefopRet)
			if mErr != nil {
				fmt.Println("parse prefop error", mErr)
				return
			}

			inputBucket := prefopRet.InputBucket
			inputKey := prefopRet.InputKey
			inputPipe := strings.Split(prefopRet.Pipeline, ".")[1]
			inputFops := make([]string, 0, len(prefopRet.Items))
			for _, item := range prefopRet.Items {
				inputFops = append(inputFops, item.Cmd)
			}

			pfopResult := PfopResult{}

			pfopParams := map[string][]string{
				"bucket": []string{inputBucket},
				"key":    []string{inputKey},
				"fops":   []string{strings.Join(inputFops, ";")},
			}
			if NotifyUrl != "" {
				pfopParams["notifyURL"] = []string{NotifyUrl}
			}
			if inputPipe != "" {
				pfopParams["pipeline"] = []string{inputPipe}
			}

			pfopParams["force"] = []string{"1"}

			err := rpcClient.CallWithForm(nil, &pfopResult, PfopUrl, pfopParams)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(pid + "\t" + pfopResult.PersistentId)
			}
		}()
	}
}
