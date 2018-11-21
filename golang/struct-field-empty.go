package main

import (
	"fmt"
	"reflect"
)

type Member struct {
	Name    string
	Age     int
	Address *Address
}

type Address struct {
	Name string
}

func collectNonEmptyFields(v interface{}) (fields []string) {
	fields = make([]string, 0)
	val := reflect.ValueOf(v)
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldName := val.Type().Field(i).Name
		fieldKind := fieldVal.Kind()
		switch fieldKind {
		case reflect.Struct, reflect.Ptr:
			if !fieldVal.IsNil() {
				fields = append(fields, fieldName)
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if fieldVal.Int() != 0 {
				fields = append(fields, fieldName)
			}
		case reflect.Bool:
			fields = append(fields, fieldName)
		case reflect.String:
			if fieldVal.String() != "" {
				fields = append(fields, fieldName)
			}
		}
	}
	return
}

func main() {
	a := Member{
		Name: "jemy",
		Age:  28,
		Address: &Address{
			Name: "Shanghai, China",
		},
	}
	fields := collectNonEmptyFields(a)
	for _, f := range fields {
		fmt.Println(f)
	}
}
