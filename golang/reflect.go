package main

import (
	"fmt"
	"reflect"
)

type Dog struct {
	Name string
	Age  int
}

func main() {
	var d = Dog{
		Name: "jemy",
		Age:  28,
	}

	dogVal := reflect.ValueOf(d)
	fieldsNum := dogVal.NumField()
	fmt.Println(fieldsNum)

	for i := 0; i < fieldsNum; i++ {
		k := reflect.Value
		fmt.Println(k)
	}
}
