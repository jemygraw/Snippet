package main

import (
	"fmt"
	"math/rand"
	"time"
)

var AlphaDigits = "abcdefghijklmnopqrstuvwxyz0123456789"

func createRandomStr(size int) string {
	alphaDigitsLen := len(AlphaDigits)
	rand.Seed(time.Now().UnixNano())
	slice := make([]byte, size)
	for i := 0; i < size; i++ {
		slice[i] = AlphaDigits[rand.Intn(alphaDigitsLen)]
	}
	return string(slice)
}

func main() {
	fmt.Println(createRandomStr(32))
}
