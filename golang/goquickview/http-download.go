package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// download the oschina home page
func main() {
	oschinaLink := "http://www.oschina.net/"

	// get home page
	respBody, readErr := http.Get(oschinaLink)
	if readErr != nil {
		fmt.Println("Err: wow!!!", readErr)
		return
	}

	defer respBody.Body.Close()

	// write to local file
	tmpFile, openErr := os.Create("/tmp/oschina.html")
	if openErr != nil {
		fmt.Println("Err: wow!!!", openErr)
		return
	}

	_, cpErr := io.Copy(tmpFile, respBody.Body)
	if cpErr != nil {
		fmt.Println("Err: wow!!!", cpErr)
		return
	}
}
