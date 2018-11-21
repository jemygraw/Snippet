package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//encode a string
	var str string = "金鑫鑫"
	var str_2_utf8_bs []byte = make([]byte, 0)
	for _, r := range str {
		buf := make([]byte, utf8.RuneLen(r))
		utf8.EncodeRune(buf, r)
		str_2_utf8_bs = append(str_2_utf8_bs, buf...)
	}

	var str_2_utf8_string=string(str_out)
	fmt.Println(str_2_utf8_string)

	//decode a string
	
}
