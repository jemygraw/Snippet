package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type Data struct {
	SignIPs string
}

func main() {
	file := "/Users/jemy/Developer/KeDotCom/k8sdeploy/config_templates/pki/kubernetes-csr.json"
	tpl, pErr := template.ParseFiles(file)
	if pErr != nil {
		fmt.Println(pErr)
		return
	}
	data := Data{
		SignIPs: "xxx",
	}
	buffer := bytes.NewBuffer(nil)
	pErr = tpl.Execute(buffer, data)
	if pErr != nil {
		fmt.Println(pErr)
		return
	}
	fmt.Println(string(buffer.Bytes()))
}
