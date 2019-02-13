package main

import (
	"encoding/xml"
	"fmt"
)

type Message struct {
	xml.Name `xml:"Envelope"`
	Body     Body `xml:"Body"`
}
type Body struct {
	Response Response `xml:"Response"`
}

type Response struct {
	Result string `xml:"Result"`
}

func main() {
	var message = `<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:soap="http://www.w3.org/2003/05/soap-envelope" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
	 <soap:Body>
	  <Response xmlns="http://tempuri.org/">
	   <Result>{"errcode":0,"errmsg":"ok","invaliduser":""}</Result>
	  </Response>
	 </soap:Body>
	</soap:Envelope>`

	var msgObj Message
	err := xml.Unmarshal([]byte(message), &msgObj)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msgObj.Body.Response.Result)

}
