package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const jsonStream = `
{"Name": "Ed", "Text": "Knock knock."}
{"Name": "Sam", "Text": "Who's there?"}
{"Name": "Ed", "Text": "Go fmt."}
{"Name": "Sam", "Text": "Go fmt who?"}
{"Name": "Ed", "Text": "Go fmt yourself!"}
`

type Message struct {
	Name, Text string
}

func main() {
	decoder := json.NewDecoder(bytes.NewReader([]byte(jsonStream)))

	for {
		var msg Message
		err := decoder.Decode(&msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(msg.Name, msg.Text)
	}
}
