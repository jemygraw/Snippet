package main

import (
	"fmt"
	"regexp"
)

func main() {
	rule := `^/upload/photo/(\w+)/(\w+)/(\w+)/(70|60|48|32|16|80|150)_(.*)$`
	str := "/upload/photo/0/0/0/70_u57313833790062233714.jpg"
	r := regexp.MustCompile(rule)
	fmt.Println(r.MatchString(str))
	m := r.FindAllStringSubmatch(str, -1)

	fmt.Println(m)
	fmt.Println(r.FindAllString(str, -1))
}
