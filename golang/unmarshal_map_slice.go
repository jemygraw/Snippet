package main

import (
	"encoding/json"
	"fmt"
)

var treeDataMap = []map[string][]string{
	map[string][]string{"A": []string{"B", "C"}},
	map[string][]string{"B": []string{"D", "E"}},
	map[string][]string{"C": []string{"F", "G", "H"}},
	map[string][]string{"F": []string{"J", "K", "L"}},
	map[string][]string{"H": []string{"I"}},
}

func main() {
	treeData, _ := json.Marshal(&treeDataMap)
	fmt.Println(string(treeData))

	json.Unmarshal(treeData, &treeDataMap)
	fmt.Println(treeDataMap)
	fmt.Println(treeDataMap[2])
}
