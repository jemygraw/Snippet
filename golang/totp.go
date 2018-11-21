package main

import (
	"fmt"
	"time"

	"github.com/dgryski/dgoogauth"
)

func main() {
	secret := "7D46C89053EFC00F08AADABEC340D8B6"
	time := time.Now().Unix()
	time = 1538026882507
	val := dgoogauth.ComputeCode(secret, time)
	fmt.Println(val)
}
