package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name      string
	Age       int
	Gender    string
	HobbyList []string
	Scores    map[string]int
}

func main() {
	jemy := Student{
		Name:      "金茧",
		Age:       29,
		Gender:    "男",
		HobbyList: []string{"吃饭", "睡觉", "工作", "玩耍"},
		Scores: map[string]int{
			"语文": 110,
			"数学": 120,
			"英语": 130,
		},
	}

	dataOutput, err := json.Marshal(&jemy)
	if err != nil {
		fmt.Println("Err: wow!!!", err)
		return
	}
	fmt.Println(string(dataOutput))
}
