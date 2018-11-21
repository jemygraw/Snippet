package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	return fmt.Sprintf("name: %s, age: %d", s.Name, s.Age)
}

func Enroll(stu *Student) {
	stu.Name = "enroll_" + stu.Name
	stu.Age += 1
}

func main() {
	stu := Student{
		Name: "jemy",
		Age:  28,
	}
	fmt.Println("before enroll:", stu)
	Enroll(&stu)
	fmt.Println("after enroll:", stu)
}
