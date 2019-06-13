package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	filePath := "/tmp/jinxinxin.txt"
	fileData := `{"Name":"金茧","Age":29,"Gender":"男","HobbyList":["吃饭","睡觉","工作","玩耍"],"Scores":{"数学":120,"英语":130,"语文":110}}`
	wErr := ioutil.WriteFile(filePath, []byte(fileData), 0644)
	if wErr != nil {
		fmt.Println("Err: wow!!!", wErr)
		return
	}
}
