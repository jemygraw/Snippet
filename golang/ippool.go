package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os/exec"
	"regexp"
	"time"
)

func getIpInfo(ip string) (ipInfo string, err error) {
	cmd := exec.Command("qip", ip)
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return
	}
	ipInfo = out.String()
	return
}

func main() {
	var port int
	flag.IntVar(&port, "port", 9090, "port number to listen")
	flag.Parse()

	http.HandleFunc("/get/ipinfo", func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		ip := req.FormValue("ip")
		if ip == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if match, _ := regexp.MatchString(`\d+.\d+.\d+.\d+`, ip); !match {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ipInfo, gErr := getIpInfo(ip)
		if gErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(ipInfo))
		fmt.Println(ip, "->", time.Now().Sub(start).String())
	})

	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
