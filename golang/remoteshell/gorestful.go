package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
)

func handleFunc(req *restful.Request, resp *restful.Response) {
	fmt.Println("...")
}
func main() {
	var host string
	var port int
	flag.StringVar(&host, "host", "0.0.0.0", "host to listen")
	flag.IntVar(&port, "port", 9001, "port to listen")
	flag.Parse()

	webService := new(restful.WebService)
	webService.Path("/remote/shell")
	webService.Route(webService.POST("").To(handleFunc))

	webSocket := new(restful.CompressingResponseWriter)

	restful.Add(webService)
	restful.Add(webSocket)
	//listen
	endPoint := fmt.Sprintf("%s:%d", host, port)
	err := http.ListenAndServe(endPoint, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
