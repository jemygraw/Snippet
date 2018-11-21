package main

import (
	"fmt"
)

type Dog struct {
	name string
	age  int
}

func (d *Dog) getLegs() int {
	return 4
}

func (d *Dog) dogInfo() {
	fmt.Println(d.age, d.name, d.getLegs())
}
func main() {
	d := Dog{name: "Jemy", age: 25}
	d.dogInfo()
}
