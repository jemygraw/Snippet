package main

import (
	"encoding/json"
	"fmt"
)

type FileInfo struct {
	Key      string `json:"key"`
	Fsize    int64  `json:"fsize"`
	Hash     string `json:"hash"`
	MimeType string `json:"mimeType"`
	PutTime  int64  `json:"putTime"`
	Type     int    `json:"type"`
	Status   int    `json:"status"`
}

type ListResp struct {
	Item   FileInfo `json:"item"`
	Marker string   `json:"marker"`
	Dir    string   `json:"dir"`
}

var respStr = `{"item":{"key":"files/D/93.log","hash":"FhWqsP2Lk36zuwGEFpPzXct12i-v","fsize":6,"mimeType":"application/octet-stream","putTime":15148601207652690,"type":0,"status":0},"marker":"eyJjIjowLCJrIjoiZmlsZXMvRC85My5sb2cifQ==","dir":""}
{"item":null,"marker":"eyJjIjowLCJrIjoiZmlsZXMvRC85NC5sb2cifQ==","dir":""}
{"item":null,"marker":"eyJjIjowLCJrIjoiZmlsZXMvIn0=","dir":"files/"}
{"item":null,"marker":"","dir":"mobile/"}`

func main() {
	var resp []ListResp
	err := json.Unmarshal([]byte(respStr), &resp)
	fmt.Println(err)
}
