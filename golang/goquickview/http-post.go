package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

// https://httpbin.org/
func main() {
	postURL := "https://httpbin.org/post"
	postBody := `{"Name":"金茧","Age":29,"Gender":"男","HobbyList":["吃饭","睡觉","工作","玩耍"],"Scores":{"数学":120,"英语":130,"语文":110}}`

	respBody, respErr := http.Post(postURL, "application/json", bytes.NewBuffer([]byte(postBody)))
	if respErr != nil {
		fmt.Println("Err: wow!!!", respErr)
		return
	}

	defer respBody.Body.Close()

	io.Copy(os.Stdout, respBody.Body)
}
