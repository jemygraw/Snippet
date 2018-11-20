package main

import (
	"os"

	"net/http"
)

func main() {

	serv := http.FileServer(http.Dir("."))

}
