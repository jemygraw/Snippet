package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var certFile = "/Users/jemy/Worklab/echo/kubernetes.pem"
var keyFile = "/Users/jemy/Worklab/echo/kubernetes-key.pem"

func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "albert.apple.com", "host to listen")
	flag.IntVar(&port, "port", 443, "port to listen")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()

		buffer := bytes.NewBuffer(nil)
		//print request method and URI
		buffer.WriteString(fmt.Sprintln(req.Method, req.RequestURI))

		buffer.WriteString(fmt.Sprintln())
		//print headers
		for k, v := range req.Header {
			buffer.WriteString(fmt.Sprintln(k+":", strings.Join(v, ";")))
		}
		buffer.WriteString(fmt.Sprintln())

		//print request body
		reqBody, readErr := ioutil.ReadAll(req.Body)
		if readErr != nil {
			buffer.WriteString(fmt.Sprintln(readErr))
		} else {
			buffer.WriteString(fmt.Sprintln(string(reqBody)))
		}

		fmt.Println(string(buffer.Bytes()))
		w.Write(buffer.Bytes())
	})

	//listen and serve
	err := http.ListenAndServeTLS(fmt.Sprintf("%s:%d", host, port), certFile, keyFile, nil)
	fmt.Println("listen err,", err)
}
