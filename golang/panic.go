package main

import (
	"fmt"
)

func deferTest() string {
	var name string
	defer func() {
		name = "first defer"
		fmt.Println("first set: ", name)
	}()

	defer func() {
		name = "second defer"
		fmt.Println("second set:", name)
	}()

	name = "func result"
	fmt.Println("func set:", name)

	panic("lalala~~~~")
	return name
}

func main() {
	fmt.Println(deferTest())
}
